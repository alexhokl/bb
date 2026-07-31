package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
)

type createSSHKeyOptions struct {
	key   string
	label string
}

var createSSHKeyOpts createSSHKeyOptions

func (opts *createSSHKeyOptions) validate() error {
	if strings.TrimSpace(opts.key) == "" {
		return fmt.Errorf("key is required (provide via --key or piped stdin)")
	}
	return nil
}

var createSSHKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Add a SSH public key to the current user",
	RunE:  runCreateSSHKey,
}

func init() {
	createCmd.AddCommand(createSSHKeyCmd)
	flags := createSSHKeyCmd.Flags()
	flags.StringVarP(&createSSHKeyOpts.key, "key", "k", "", "SSH public key in OpenSSH format (otherwise read from stdin when piped)")
	flags.StringVarP(&createSSHKeyOpts.label, "label", "l", "", "User-defined label for the key")
}

func runCreateSSHKey(_ *cobra.Command, _ []string) error {
	if createSSHKeyOpts.key == "" {
		stdin, err := readStdinIfPiped()
		if err != nil {
			return err
		}
		createSSHKeyOpts.key = stdin
	}
	if err := createSSHKeyOpts.validate(); err != nil {
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

	body := swagger.SshAccountKey{
		Key:   createSSHKeyOpts.key,
		Label: createSSHKeyOpts.label,
	}
	key, _, err := client.SshApi.UsersSelectedUserSshKeysPost(
		auth,
		userUUID,
		&swagger.SshApiUsersSelectedUserSshKeysPostOpts{Body: optional.NewInterface(body)},
	)
	if err != nil {
		return err
	}

	models.PrintSshAccountKey(&key)
	return nil
}
