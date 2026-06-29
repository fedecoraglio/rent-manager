package domain

import "time"

type Tenant struct {
	ID             int64
	CountryID      int64
	StateID        int64
	Name           string
	Email          string
	Phone          string
	DocumentNumber string

	City         string
	Street       string
	StreetNumber string
	Floor        string
	Apartment    string
	PostalCode   string

	CreatedAt time.Time
	UpdatedAt time.Time
}
