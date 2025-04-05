package user

import (
	"context"
	"time"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

type userRepo interface {
	findById(id string) (*user, error)
	
}

type workosUserRepo struct {
	client *usermanagement.Client
}

func newWorkosUserRepo(client *usermanagement.Client) userRepo {
	return &workosUserRepo{client: client}
}

func (r *workosUserRepo) findById(id string) (*user, error) {
	userData, err := r.client.GetUser(context.Background(), usermanagement.GetUserOpts{
		User: id,
	})
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, userData.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user{
		ID:        userData.ID,
		Email:     userData.Email,
		CreatedAt: createdAt,
	}, nil
}

var _ userRepo = &workosUserRepo{}
