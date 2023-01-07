package command

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexhokl/bb/client"
	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/helper/git"
	"github.com/alexhokl/helper/regexhelper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cli interface
type Cli interface {
	Client() client.APIClient
	UserCredential() *models.UserCredential
	Repo() *models.Repository
}

// ManagerCli struct
type ManagerCli struct {
	client     client.APIClient
	credential *models.UserCredential
	repo       *models.Repository
}

type idOption struct {
	id int
}

type idOptions struct {
	id int32
}

var idOpts idOptions

func (opts idOptions) validate() error {
	if opts.id <= 0 {
		return fmt.Errorf("invalid pull request ID")
	}

	return nil
}

// NewManagerCli creates a new manager cli instance
func NewManagerCli() *ManagerCli {
	cli := ManagerCli{
		client: client.NewClient(),
	}
	return &cli
}

// Client returns API client
func (cli *ManagerCli) Client() client.APIClient {
	return cli.client
}

// UserCredential returns user credential
func (cli *ManagerCli) UserCredential() *models.UserCredential {
	return cli.credential
}

// Repo returns information on Repository
func (cli *ManagerCli) Repo() *models.Repository {
	return cli.repo
}

// ShowHelp shows the command help
func (cli *ManagerCli) ShowHelp(cmd *cobra.Command, args []string) error {
	cmd.HelpFunc()(cmd, args)
	return nil
}

// SetCredentials retrieves credentials (access token) from a local configuration file
func (cli *ManagerCli) SetCredentials() error {
	accessToken := viper.GetString("access_token")
	if accessToken == "" {
		return fmt.Errorf("please run command login before continue on")
	}
	refreshToken := viper.GetString("refresh_token")

	if cli.credential == nil {
		cli.credential = &models.UserCredential{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		return nil
	}
	cli.credential.AccessToken = accessToken
	cli.credential.RefreshToken = refreshToken

	return nil
}

// SetJiraCredentials retrieves JIRA credentials from either configuration file
// or environment variables
func (cli *ManagerCli) SetJiraCredentials() error {
	email := viper.GetString("jira_email")
	if email == "" {
		return fmt.Errorf("email address of JIRA has not been configured")
	}
	key := viper.GetString("jira_api_key")
	if key == "" {
		return fmt.Errorf("API Key of JIRA has not been configured")
	}

	if cli.credential == nil {
		cli.credential = &models.UserCredential{
			JiraEmailAddress: email,
			JiraAPIKey:       key,
		}
		return nil
	}
	cli.credential.JiraEmailAddress = email
	cli.credential.JiraAPIKey = key
	return nil
}

// SetRepository sets repository information using information from git setup
// in the current directory
func (cli *ManagerCli) SetRepository() error {
	cmdOut, err := git.GetOriginURL()
	if err != nil {
		if strings.Contains(err.Error(), "129") {
			return errors.New("git remote get-url is not supported. Please upgrade to the latest version of git")
		}
		return err
	}

	remote := string(cmdOut)
	if !strings.Contains(remote, "bitbucket.org") {
		return errors.New("Error: Only repository of BitBucket is supported")
	}

	regex := regexp.MustCompile(`bitbucket\.org/(?P<org>\w+)\/(?P<name>.*)`)
	matches := regexhelper.FindNamedGroupMatchedStrings(regex, remote)

	if matches["org"] == "" || matches["name"] == "" {
		return fmt.Errorf("Error: Unable to retrieve repository name")
	}

	cli.repo = &models.Repository{
		Org:  matches["org"],
		Name: matches["name"],
	}

	return nil
}
