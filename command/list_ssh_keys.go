package command

import (
	"context"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
)

var listSSHKeysCmd = &cobra.Command{
	Use:     "ssh-keys",
	Aliases: []string{"ssh-key"},
	Short:   "List SSH keys of the current user",
	RunE:    runListSSHKeys,
}

func init() {
	listCmd.AddCommand(listSSHKeysCmd)
}

func runListSSHKeys(_ *cobra.Command, _ []string) error {
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

	page, _, err := client.SshApi.UsersSelectedUserSshKeysGet(auth, userUUID)
	if err != nil {
		return err
	}

	for i := range page.Values {
		models.PrintSshAccountKey(&page.Values[i])
	}
	return nil
}
