package user

import "time"

type User struct {
	ID              int
	Name            string
	Email           string
	PasswordHash    string
	Occupation      string
	Avatar_filename string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
