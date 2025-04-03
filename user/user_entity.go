package user

import "time"

type User struct {
	ID        string
	Email     string
	CreatedAt time.Time
}

func (u *User) dto() UserDTO {
	return UserDTO{
		ID:        u.ID,
		Email:     u.Email,
	}
}
