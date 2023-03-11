package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a subject of subcommand",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please specify the subject to be deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

