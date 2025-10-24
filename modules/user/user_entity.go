package user

import "time"

type user struct {
	ID        string
	Email     string
	CreatedAt time.Time
}

func (u *user) dto() UserDTO {
	return UserDTO{
		ID:    u.ID,
		Email: u.Email,
	}
}
