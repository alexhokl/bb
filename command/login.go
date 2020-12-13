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
	"os/user"
	"time"

	jsonhelper "github.com/alexhokl/helper/json"
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
const tokenDirectory = ".go_bb_pr"
const tokenFilename = "token.json"

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

// GetTokenPath returns full path to token JSON file
func GetTokenPath() (string, error) {
	directory, err := getTokenDirectory(tokenDirectory)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", directory, tokenFilename), nil
}

func runLogin(cli *ManagerCli, opts loginOptions) error {
	clientID := viper.GetString("client_id")
	clientSecret := viper.GetString("client_secret")

	if clientID == "" || clientSecret == "" {
		return fmt.Errorf("client_id or client_secret is not configured")
	}

	tokenDirectoryFullPath, errDir := getTokenDirectory(tokenDirectory)
	if errDir != nil {
		return fmt.Errorf("Unable to create token directory [$HOME/%s]: %v", tokenDirectory, errDir)
	}
	errEnsureDir := ensureDirectory(tokenDirectoryFullPath)
	if errEnsureDir != nil {
		return fmt.Errorf("Unable to create token directory [%s]: %v", tokenDirectoryFullPath, errEnsureDir)
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
	cmdName, cmdArgs := getOpenCommand(url)
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

	tokenPath, _ := GetTokenPath()
	errJSON := jsonhelper.WriteToJSONFile(tokenPath, token, true)
	if errJSON != nil {
		fmt.Printf("Unable to write token to [%s]: %v", tokenPath, errJSON)
		os.Exit(1)
		return
	}

	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
	fmt.Fprintf(w, msg)

	fmt.Printf("Login has been completed successfully. Tokens are stored in [%s]\n", tokenPath)

	go func() {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()
}

func getTokenDirectory(path string) (string, error) {
	usr, errCurrent := user.Current()
	if errCurrent != nil {
		return "", errCurrent
	}
	return fmt.Sprintf("%s/%s", usr.HomeDir, path), nil
}

func ensureDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		errMkdir := os.Mkdir(path, 0755)
		if errMkdir != nil {
			return errMkdir
		}
	}
	return nil
}
