package domain

import "time"

type InflationIndex struct {
	ID         int64
	Period     time.Time
	Percentage float64
	Source     string
	Notes      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
