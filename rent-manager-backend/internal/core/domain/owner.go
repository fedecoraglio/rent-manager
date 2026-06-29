package domain

import "time"

type Owner struct {
	ID             int64
	Name           string
	Email          string
	Phone          string
	DocumentNumber string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
