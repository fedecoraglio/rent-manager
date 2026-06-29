package domain

import "time"

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	RoleID       int64
	Role         *Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
