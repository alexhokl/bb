package command

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"

	clihelper "github.com/alexhokl/helper/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/bitbucket"
)

type loginOptions struct {
	isDebug bool
}

var (
	conf *oauth2.Config
	ctx  context.Context
)

const port = 9990

// NewLoginCommand returns definition of command list
func NewLoginCommand(cli *ManagerCli) *cobra.Command {
	opts := loginOptions{}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Log in to BitBucket",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runLogin(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.isDebug, "debug", "d", false, "Debug mode")

	return cmd
}

func runLogin(cli *ManagerCli, opts loginOptions) error {
	clientID := viper.GetString("client_id")
	clientSecret := viper.GetString("client_secret")

	if clientID == "" || clientSecret == "" {
		return fmt.Errorf("client_id or client_secret is not configured")
	}

	ctx = context.Background()
	conf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{}, // scopes are configured on OAuth Consumers on BitBucket
		Endpoint:     bitbucket.Endpoint,
		RedirectURL:  fmt.Sprintf("http://localhost:%d/oauth/callback", port),
	}

	// add transport for self-signed certificate to context
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Println("You will now be taken to your browser for authentication")
	time.Sleep(1 * time.Second)
	cmdName, cmdArgs := clihelper.GetOpenCommand(url)
	_, errOpen := exec.Command(cmdName, cmdArgs...).Output()
	if errOpen != nil {
		return errOpen
	}
	time.Sleep(1 * time.Second)

	http.HandleFunc("/oauth/callback", callbackHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

	return nil
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)
	code := queryParts["code"][0]

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	viper.Set("access_token", token.AccessToken)
	viper.Set("refresh_token", token.RefreshToken)
	viper.WriteConfig()

	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
	fmt.Fprintf(w, msg)

	fmt.Println("Login has been completed successfully.")

	go func() {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()
}
