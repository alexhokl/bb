package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/alexhokl/go-bb-pr/command"
	"github.com/alexhokl/go-bb-pr/git"
	"github.com/alexhokl/go-bb-pr/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func main() {
	viper.SetEnvPrefix("bb_pr")
	viper.AutomaticEnv()

	cred, errCred := getCredentials()
	if errCred != nil {
		fmt.Println(errCred)
		os.Exit(1)
	}

	repo, errRemote := getRepository()
	if errRemote != nil {
		fmt.Println(errRemote)
		os.Exit(1)
	}

	managerCli := command.NewManagerCli(cred, repo)
	cmd := newManagerCommand(managerCli)

	if err := cmd.Execute(); err != nil {
		if sterr, ok := err.(command.StatusError); ok {
			if sterr.Status != "" {
				fmt.Println(sterr.Status)
			}
			if sterr.StatusCode == 0 {
				os.Exit(1)
			}
			os.Exit(sterr.StatusCode)
		}
		os.Exit(1)
	}
}

func newManagerCommand(cli *command.ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "go-bb-pr",
		Short:        "BitBucket Pull Request Manager",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.ShowHelp(cmd, args)
		},
	}
	command.AddCommands(cmd, cli)
	return cmd
}

func getCredentials() (*models.UserCredential, error) {
	tokenPath, errPath := command.GetTokenPath()
	if errPath != nil {
		return nil, errPath
	}
	if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Please run command login before continue on")
	}

	file, errFile := ioutil.ReadFile(tokenPath)
	if errFile != nil {
		return nil, fmt.Errorf("Please run command login before continue on: %v", errFile)
	}

	token := oauth2.Token{}
	err := json.Unmarshal(file, &token)
	if err != nil {
		return nil, fmt.Errorf("Please run command login before continue on: %v", err)
	}

	cred := &models.UserCredential{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return cred, nil
}

func getRepository() (*models.Repository, error) {
	cmdOut, err := git.GetOriginURL()
	if err != nil {
		if strings.Contains(err.Error(), "129") {
			return nil, errors.New("git remote get-url is not supported. Please upgrade to the latest version of git")
		}
		return nil, err
	}

	remote := string(cmdOut)
	if !strings.Contains(remote, "bitbucket.org") {
		return nil, errors.New("Error: Only repository of BitBucket is supported")
	}

	regex := regexp.MustCompile(`bitbucket\.org/(?P<org>\w+)\/(?P<name>.*)`)
	matches := findMatches(regex, remote)

	if matches["org"] == "" || matches["name"] == "" {
		return nil, fmt.Errorf("Error: Unable to retrieve repository name")
	}

	r := models.Repository{
		Org:  matches["org"],
		Name: matches["name"],
	}

	return &r, nil
}

func findMatches(regex *regexp.Regexp, input string) map[string]string {
	match := regex.FindStringSubmatch(input)
	subMatchMap := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}

	return subMatchMap
}
