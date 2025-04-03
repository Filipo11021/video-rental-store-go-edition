package user

type UserDTO struct {
	ID        string
	Email     string
}


type VerifyMagicLinkParamsDTO struct {
	Code string
	Email string
}