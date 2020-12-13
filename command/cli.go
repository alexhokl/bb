package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/alexhokl/go-bb-pr/client"
	"github.com/alexhokl/go-bb-pr/models"
	"github.com/alexhokl/helper/git"
	"github.com/alexhokl/helper/iohelper"
	"github.com/alexhokl/helper/regexhelper"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
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
	tokenPath, errPath := getTokenPath()
	if errPath != nil {
		return errPath
	}
	if !iohelper.IsFileExist(tokenPath) {
		return fmt.Errorf("Please run command login before continue on")
	}

	file, errFile := ioutil.ReadFile(tokenPath)
	if errFile != nil {
		return fmt.Errorf("Please run command login before continue on: %v", errFile)
	}

	token := oauth2.Token{}
	err := json.Unmarshal(file, &token)
	if err != nil {
		return fmt.Errorf("Please run command login before continue on: %v", err)
	}

	cli.credential = &models.UserCredential{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

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
