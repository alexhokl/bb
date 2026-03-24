package command

import (
	"fmt"

	"github.com/alexhokl/bb/models"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

type getUserOptions struct {
	accountID string
}

var getUserOpts getUserOptions

func (opts *getUserOptions) validate() error {
	if opts.accountID == "" {
		return fmt.Errorf("account ID is required")
	}
	return nil
}

var getUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Get a user by account ID or UUID",
	RunE:  runGetUser,
}

func init() {
	getCmd.AddCommand(getUserCmd)
	flags := getUserCmd.Flags()
	flags.StringVarP(&getUserOpts.accountID, "account", "a", "", "Account ID or UUID of the user")
	_ = getUserCmd.MarkFlagRequired("account")
}

func runGetUser(_ *cobra.Command, _ []string) error {
	if err := getUserOpts.validate(); err != nil {
		return err
	}

	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	config := swagger.NewConfiguration()
	client := swagger.NewAPIClient(config)

	account, _, err := client.UsersApi.UsersSelectedUserGet(auth, getUserOpts.accountID)
	if err != nil {
		return err
	}

	models.PrintAccount(&account)
	return nil
}
