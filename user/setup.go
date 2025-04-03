package user

import (
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"go.uber.org/fx"
)

type CreateConfigDTO struct {
	WorkOSAPIKey   string
	WorkOSClientID string
}

type config struct {
	workOSAPIKey   string
	workOSClientID string
}

func Module(c CreateConfigDTO) fx.Option {
	client := usermanagement.NewClient(c.WorkOSAPIKey)

	return fx.Module("user",
		fx.Provide(
			newFacade,
		),
		fx.Provide(
			fx.Private,
			func() *usermanagement.Client {
				return client
			},
			newWorkosUserRepo,
			func() magicLinkRepo {
				return newWorkOSMagicLinkRepo(client, c.WorkOSClientID)
			},
		),
	)
}
