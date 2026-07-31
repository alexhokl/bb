package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type deleteGpgKeyOptions struct {
	fingerprint string
}

var deleteGpgKeyOpts deleteGpgKeyOptions

func (opts *deleteGpgKeyOptions) validate() error {
	if strings.TrimSpace(opts.fingerprint) == "" {
		return fmt.Errorf("fingerprint is required (pass --fingerprint)")
	}
	return nil
}

var deleteGpgKeyCmd = &cobra.Command{
	Use:   "gpg-key",
	Short: "Delete a GPG public key from the current user",
	RunE:  runDeleteGpgKey,
}

func init() {
	deleteCmd.AddCommand(deleteGpgKeyCmd)
	flags := deleteGpgKeyCmd.Flags()
	flags.StringVar(&deleteGpgKeyOpts.fingerprint, "fingerprint", "", "Fingerprint of the GPG key (see bb list gpg-keys)")
	_ = deleteGpgKeyCmd.MarkFlagRequired("fingerprint")
}

func runDeleteGpgKey(_ *cobra.Command, _ []string) error {
	if err := deleteGpgKeyOpts.validate(); err != nil {
		return err
	}

	client, err := newGpgClient()
	if err != nil {
		return err
	}

	if err := client.Delete(deleteGpgKeyOpts.fingerprint); err != nil {
		return err
	}

	fmt.Printf("GPG key %s deleted\n", deleteGpgKeyOpts.fingerprint)
	return nil
}
