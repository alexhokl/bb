package main

import (
	"fmt"
	"os"

	"github.com/alexhokl/bb/command"
	"github.com/alexhokl/helper/cli"
	"github.com/spf13/cobra"
)

var configurationFilePath string
var verbose bool

func main() {
	managerCli := command.NewManagerCli()
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
		Use:          "bb",
		Short:        "BitBucket Pull Request Manager",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.ShowHelp(cmd, args)
		},
	}

	cmd.PersistentFlags().StringVar(&configurationFilePath, "config", "", "config file (default is $HOME/.bb_pr.yaml)")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode (default is off)")

	cobra.OnInitialize(initConfig)

	command.AddCommands(cmd, cli)
	return cmd
}

func initConfig() {
	cli.ConfigureViper(configurationFilePath, "bb", verbose, "bb")
}
