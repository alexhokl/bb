package models

import (
	"fmt"

	"github.com/alexhokl/bb/internal/gpgapi"
	"github.com/fatih/color"
)

// PrintGpgAccountKey prints a single GPG public key entry.
func PrintGpgAccountKey(k *gpgapi.GpgAccountKey) {
	color.New(color.FgCyan).Printf("%s", k.Fingerprint)
	if k.KeyID != "" {
		fmt.Printf(" %s", k.KeyID)
	}
	if k.Name != "" {
		color.New(color.FgGreen).Printf(" [%s]", k.Name)
	}
	fmt.Println()
}
