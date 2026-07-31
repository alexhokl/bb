package models

import (
	"fmt"

	"github.com/alexhokl/bb/swagger"
	"github.com/fatih/color"
)

// PrintSshAccountKey prints a single SSH public key entry.
func PrintSshAccountKey(k *swagger.SshAccountKey) {
	color.New(color.FgCyan).Printf("%s", k.Uuid)
	if k.Comment != "" {
		fmt.Printf(" %s", k.Comment)
	}
	if k.Label != "" {
		color.New(color.FgGreen).Printf(" [%s]", k.Label)
	}
	fmt.Println()
}
