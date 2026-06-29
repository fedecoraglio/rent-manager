package domain

import "time"

type Role struct {
	ID        int64
	Code      string
	Name      string
	CreatedAt time.Time
}
