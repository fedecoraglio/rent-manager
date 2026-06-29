package domain

import "time"

type TokenClaims struct {
	UserID    int64
	Email     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}
