package user

type Api interface {
	SendMagicLink(email string) error
	VerifyMagicLink(dto VerifyMagicLinkParamsDTO) error
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

func (a *api) VerifyMagicLink(dto VerifyMagicLinkParamsDTO) error {
	err := a.magicLinkRepo.verifyMagicLink(dto)
	return err
}

func (a *api) GetUserById(id string) (*UserDTO, error) {
	user, err := a.userRepo.findById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	userDTO := user.dto()
	return &userDTO, nil
}
