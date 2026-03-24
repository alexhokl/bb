package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Update a project",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please specify the action to be performed on the project")
	},
}

func init() {
	updateCmd.AddCommand(updateProjectCmd)
}
