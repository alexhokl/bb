package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/alexhokl/go-bb-pr/command"
	"github.com/alexhokl/go-bb-pr/git"
	"github.com/alexhokl/go-bb-pr/models"
	"github.com/spf13/cobra"
)

func main() {
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
		Use:          "pr",
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
	username := os.Getenv("bbuser")
	password := os.Getenv("bbpassword")

	if username == "" {
		return nil, errors.New("bbuser is not set")
	}
	if password == "" {
		return nil, errors.New("bbpassword is not set")
	}

	cred := models.UserCredential{Username: username, Password: password}

	return &cred, nil
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
		return nil, errors.New("Only BitBucket remote is supported")
	}

	remote = strings.Replace(remote, "https://bitbucket.org/", "", -1)
	remote = strings.Replace(remote, ".git", "", -1)
	remote = strings.Replace(remote, "\n", "", -1)

	parts := strings.Split(remote, "/")

	r := models.Repository{
		Org:  parts[0],
		Name: parts[1],
	}

	return &r, nil
}
