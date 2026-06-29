package domain

import "time"

type State struct {
	ID        int64
	CountryID int64
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
