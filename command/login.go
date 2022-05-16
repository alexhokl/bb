package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/helper/authhelper"
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
	token, err := authhelper.GetToken(ctx, authhelper.OAuthConfig{
		ClientId:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{}, // scopes are configured on OAuth Consumers on BitBucket
		RedirectURI:  "/oauth/callback",
		Port:         port,
		Endpoint:     bitbucket.Endpoint,
	})
	if err != nil {
		return err
	}

	viper.Set("access_token", token.AccessToken)
	viper.Set("refresh_token", token.RefreshToken)
	viper.WriteConfig()

	fmt.Println("Login has been completed successfully.")

	return nil
}
