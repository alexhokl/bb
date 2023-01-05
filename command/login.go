package command

import (
	"context"
	"fmt"

	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ctx  context.Context
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log onto BitBucket",
	RunE:  runLogin,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func runLogin(_ *cobra.Command, _ []string) error {
	config, err := getOAuthConfigurationFromViper()
	if err != nil {
		return err
	}

	ctx = context.Background()
	token, err := authhelper.GetToken(ctx, config, false)
	if err != nil {
		return err
	}

	authhelper.SaveTokenToViper(token)
	viper.WriteConfig()
	fmt.Println("Login has been completed successfully")

	return nil
}

func getScopes() []string {
	// scopes are configured on OAuth Consumers on BitBucket
	return []string{}
}
