package user

type Facade interface {
	SendMagicLink(email string) error
	VerifyMagicLink(dto VerifyMagicLinkParamsDTO) error
	GetUserById(id string) (*UserDTO, error)
}

type facade struct {
	userRepo      userRepo
	magicLinkRepo magicLinkRepo
}

func newFacade(
	userRepo userRepo,
	magicLinkRepo magicLinkRepo,
) Facade {
	return &facade{
		userRepo:      userRepo,
		magicLinkRepo: magicLinkRepo,
	}
}

func (f *facade) SendMagicLink(email string) error {
	err := f.magicLinkRepo.sendMagicLink(email)
	return err
}

func (f *facade) VerifyMagicLink(dto VerifyMagicLinkParamsDTO) error {
	err := f.magicLinkRepo.verifyMagicLink(dto)
	return err
}

func (f *facade) GetUserById(id string) (*UserDTO, error) {
	user, err := f.userRepo.findById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	userDTO := user.dto()
	return &userDTO, nil
}
