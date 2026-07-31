package command

import (
	"fmt"
	"strings"

	"github.com/alexhokl/bb/models"
	"github.com/spf13/cobra"
)

type createGpgKeyOptions struct {
	key  string
	name string
}

var createGpgKeyOpts createGpgKeyOptions

func (opts *createGpgKeyOptions) validate() error {
	if strings.TrimSpace(opts.key) == "" {
		return fmt.Errorf("key is required (provide via --key or piped stdin)")
	}
	if strings.TrimSpace(opts.name) == "" {
		return fmt.Errorf("name is required (pass --name)")
	}
	return nil
}

var createGpgKeyCmd = &cobra.Command{
	Use:   "gpg-key",
	Short: "Add a GPG public key to the current user",
	RunE:  runCreateGpgKey,
}

func init() {
	createCmd.AddCommand(createGpgKeyCmd)
	flags := createGpgKeyCmd.Flags()
	flags.StringVarP(&createGpgKeyOpts.key, "key", "k", "", "ASCII-armoured GPG public key (otherwise read from stdin when piped)")
	flags.StringVarP(&createGpgKeyOpts.name, "name", "n", "", "User-defined name for the key (required)")
	_ = createGpgKeyCmd.MarkFlagRequired("name")
}

func runCreateGpgKey(_ *cobra.Command, _ []string) error {
	if createGpgKeyOpts.key == "" {
		stdin, err := readStdinIfPiped()
		if err != nil {
			return err
		}
		createGpgKeyOpts.key = stdin
	}
	if err := createGpgKeyOpts.validate(); err != nil {
		return err
	}

	client, err := newGpgClient()
	if err != nil {
		return err
	}

	key, err := client.Add(createGpgKeyOpts.name, createGpgKeyOpts.key)
	if err != nil {
		return err
	}

	models.PrintGpgAccountKey(key)
	return nil
}
