package command

import (
	"github.com/alexhokl/bb/internal/gpgapi"
	"github.com/alexhokl/bb/swagger"
	"github.com/alexhokl/helper/authhelper"
	"golang.org/x/net/context"
)

// newGpgClient returns a GPG-keys API client bound to the current user.
func newGpgClient() (*gpgapi.Client, error) {
	savedToken, err := authhelper.LoadTokenFromViper()
	if err != nil {
		return nil, err
	}
	auth := context.WithValue(context.Background(), swagger.ContextAccessToken, savedToken.AccessToken)
	swaggerClient := swagger.NewAPIClient(swagger.NewConfiguration())

	userUUID, err := getCurrentUserUUID(swaggerClient, auth)
	if err != nil {
		return nil, err
	}
	return gpgapi.New(savedToken.AccessToken, userUUID), nil
}
