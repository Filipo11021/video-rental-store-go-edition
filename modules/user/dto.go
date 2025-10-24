package user

type UserDTO struct {
	ID    string
	Email string
}

type VerifyMagicLinkParamsDto struct {
	Email string
	Code  string
}

type AuthTokensDto struct {
	AccessToken  string
	RefreshToken string
}

type AuthenticateResponseDto struct {
	AccessToken  string
	RefreshToken string
}
