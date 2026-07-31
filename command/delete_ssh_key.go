package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

type deleteSSHKeyOptions struct {
	id string
}

var deleteSSHKeyOpts deleteSSHKeyOptions

func (opts *deleteSSHKeyOptions) validate() error {
	if strings.TrimSpace(opts.id) == "" {
		return fmt.Errorf("key id is required (pass --id)")
	}
	return nil
}

var deleteSSHKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Delete a SSH public key from the current user",
	RunE:  runDeleteSSHKey,
}

func init() {
	deleteCmd.AddCommand(deleteSSHKeyCmd)
	flags := deleteSSHKeyCmd.Flags()
	flags.StringVarP(&deleteSSHKeyOpts.id, "id", "i", "", "UUID of the SSH key (see bb list ssh-keys)")
	_ = deleteSSHKeyCmd.MarkFlagRequired("id")
}

func runDeleteSSHKey(_ *cobra.Command, _ []string) error {
	if err := deleteSSHKeyOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	userUUID, err := getCurrentUserUUID(client, auth)
	if err != nil {
		return err
	}

	_, err = client.SshApi.UsersSelectedUserSshKeysKeyIdDelete(auth, deleteSSHKeyOpts.id, userUUID)
	if err != nil {
		return err
	}

	fmt.Printf("SSH key %s deleted\n", deleteSSHKeyOpts.id)
	return nil
}
