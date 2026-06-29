package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"rent-manager-backend/internal/core/domain"
)

type JWTTokenProvider struct {
	secretKey []byte
	ttl       time.Duration
}

func NewJWTTokenProvider(secretKey string, ttl time.Duration) *JWTTokenProvider {
	return &JWTTokenProvider{
		secretKey: []byte(secretKey),
		ttl:       ttl,
	}
}

func (p *JWTTokenProvider) GenerateToken(user *domain.User) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"iat":     now.Unix(),
		"exp":     now.Add(p.ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(p.secretKey)
}

func (p *JWTTokenProvider) ValidateToken(tokenValue string) (*domain.TokenClaims, error) {
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}

		return p.secretKey, nil
	})
	if err != nil {
		return nil, domain.ErrUnauthorized
	}

	if !token.Valid {
		return nil, domain.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	issuedAtFloat, ok := claims["iat"].(float64)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	expiresAtFloat, ok := claims["exp"].(float64)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	return &domain.TokenClaims{
		UserID:    int64(userIDFloat),
		Email:     email,
		IssuedAt:  time.Unix(int64(issuedAtFloat), 0),
		ExpiresAt: time.Unix(int64(expiresAtFloat), 0),
	}, nil
}
