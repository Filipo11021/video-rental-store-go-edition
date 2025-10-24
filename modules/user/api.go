package user

import "errors"

type Api interface {
	SendMagicLink(email string) error
	VerifyMagicLink(dto VerifyMagicLinkParamsDto) (AuthTokensDto, error)
	GetUserById(id string) (*UserDTO, error)
}

type api struct {
	userRepo      userRepo
	magicLinkRepo magicLinkRepo
}

func newApi(
	userRepo userRepo,
	magicLinkRepo magicLinkRepo,
) Api {
	return &api{
		userRepo:      userRepo,
		magicLinkRepo: magicLinkRepo,
	}
}

func (a *api) SendMagicLink(email string) error {
	err := a.magicLinkRepo.sendMagicLink(email)
	return err
}

func (a *api) VerifyMagicLink(dto VerifyMagicLinkParamsDto) (AuthTokensDto, error) {
	result, err := a.magicLinkRepo.verifyMagicLink(VerifyMagicLinkParamsDto{
		Email: dto.Email,
		Code:  dto.Code,
	})
	if err != nil {
		return AuthTokensDto{}, err
	}
	return result, nil
}

func (a *api) GetUserById(id string) (*UserDTO, error) {
	user, err := a.userRepo.findById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	userDTO := user.dto()
	return &userDTO, nil
}
