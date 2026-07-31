package command

import (
	"github.com/alexhokl/bb/models"
	"github.com/spf13/cobra"
)

var listGpgKeysCmd = &cobra.Command{
	Use:     "gpg-keys",
	Aliases: []string{"gpg-key"},
	Short:   "List GPG keys of the current user",
	RunE:    runListGpgKeys,
}

func init() {
	listCmd.AddCommand(listGpgKeysCmd)
}

func runListGpgKeys(_ *cobra.Command, _ []string) error {
	client, err := newGpgClient()
	if err != nil {
		return err
	}

	keys, err := client.List()
	if err != nil {
		return err
	}

	for i := range keys {
		models.PrintGpgAccountKey(&keys[i])
	}
	return nil
}
