package domain

import "time"

type PropertyType struct {
	ID   int64
	Code string
	Name string
}

type PropertyStatus struct {
	ID   int64
	Code string
	Name string
}

type Property struct {
	ID        int64
	OwnerID   int64
	TypeID    int64
	StatusID  int64
	CountryID int64
	StateID   int64

	Code        string
	Title       string
	Description string

	Street       string
	StreetNumber string
	Floor        string
	Apartment    string
	City         string
	PostalCode   string

	Owner   *Owner
	Type    *PropertyType
	Status  *PropertyStatus
	Country *Country
	State   *State

	CreatedAt time.Time
	UpdatedAt time.Time
}
