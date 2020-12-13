package main

import (
	"fmt"
	"os"

	"github.com/alexhokl/go-bb-pr/command"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("bb_pr")
	viper.AutomaticEnv()

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
