package user

import (
	"context"
	"fmt"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

type magicLinkRepo interface {
	sendMagicLink(email string) error
	verifyMagicLink(params VerifyMagicLinkParamsDto) (AuthTokensDto, error)
}

type workosMagicLinkRepo struct {
	clientId string
	client   *usermanagement.Client
}

func newWorkOSMagicLinkRepo(client *usermanagement.Client, clientId string) magicLinkRepo {
	return &workosMagicLinkRepo{
		clientId: clientId,
		client:   client,
	}
}

func (r *workosMagicLinkRepo) sendMagicLink(email string) error {
	_, err := r.client.CreateMagicAuth(
		context.Background(),
		usermanagement.CreateMagicAuthOpts{
			Email: email,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create magic auth: %w", err)
	}

	return nil
}

func (r *workosMagicLinkRepo) verifyMagicLink(params VerifyMagicLinkParamsDto) (AuthTokensDto, error) {
	result, err := r.client.AuthenticateWithMagicAuth(
		context.Background(),
		usermanagement.AuthenticateWithMagicAuthOpts{
			ClientID: r.clientId,
			Email:    params.Email,
			Code:     params.Code,
		},
	)
	if err != nil {
		return AuthTokensDto{}, fmt.Errorf("failed to get magic auth: %w", err)
	}

	return AuthTokensDto{
		AccessToken:  result.AccessToken,
		RefreshToken: result.AccessToken,
	}, nil
}
