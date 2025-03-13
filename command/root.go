package command

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/helper/authhelper"
	"github.com/alexhokl/helper/cli"
	"github.com/alexhokl/helper/git"
	"github.com/alexhokl/helper/regexhelper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/bitbucket"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:               "bb",
	Short:             "BitBucket Pull Request Manager",
	SilenceUsage:      true,
	PersistentPreRunE: validateToken,
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bb.yml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	cli.ConfigureViper(cfgFile, "bb", false, "")
}

func validateToken(cmd *cobra.Command, _ []string) error {
	// skips checking if it is login
	if cmd.Name() == "login" {
		return nil
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}

	config, err := getOAuthConfigurationFromViper()
	if err != nil {
		return err
	}

	if !savedToken.Valid() {
		ctx := context.Background()
		newToken, err := authhelper.RefreshToken(ctx, config.GetOAuthConfig(), savedToken)
		if err != nil {
			return fmt.Errorf("invalid token. please login again: %v", err)
		}
		authhelper.SaveTokenToViper(newToken)
		return nil
	}

	return nil
}

// getRepositoryInfoFromCurrentPath returns repository information using
// information from git setup in the current directory
func getRepositoryInfoFromCurrentPath() (*models.Repository, error) {
	cmdOut, err := git.GetOriginURL()
	if err != nil {
		if strings.Contains(err.Error(), "129") {
			return nil, fmt.Errorf("git remote get-url is not supported. Please upgrade to the latest version of git")
		}
		return nil, err
	}

	remote := string(cmdOut)
	if !strings.Contains(remote, "bitbucket.org") {
		return nil, fmt.Errorf("Error: Only repository of BitBucket is supported")
	}

	regex := regexp.MustCompile(`bitbucket\.org/(?P<org>\w+)\/(?P<name>.*)`)
	matches := regexhelper.FindNamedGroupMatchedStrings(regex, remote)
	if len(matches) == 0 {
		alternativeRegex := regexp.MustCompile(`bitbucket\.org:(?P<org>\w+)/(?P<name>.*)`)
		matches = regexhelper.FindNamedGroupMatchedStrings(alternativeRegex, remote)
		if len(matches) == 0 {
			return nil, fmt.Errorf("Error: Unable to retrieve repository name")
		}
	}

	if matches["org"] == "" || matches["name"] == "" {
		return nil, fmt.Errorf("Error: Unable to retrieve repository name")
	}

	repo := &models.Repository{
		Org:  matches["org"],
		Name: matches["name"],
	}

	return repo, nil
}

func getOAuthConfigurationFromViper() (*authhelper.OAuthConfig, error) {
	clientID := viper.GetString("client_id")
	clientSecret := viper.GetString("client_secret")
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("client_id or client_secret is not configured")
	}

	config := &authhelper.OAuthConfig{
		ClientId:     clientID,
		ClientSecret: clientSecret,
		Scopes:       getScopes(),
		RedirectURI:  "/oauth/callback",
		Port:         viper.GetInt("port"),
		Endpoint:     bitbucket.Endpoint,
	}
	return config, nil
}
