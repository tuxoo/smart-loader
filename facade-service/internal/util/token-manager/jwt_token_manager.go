package token_manager

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTTokenManager struct {
	ttl        time.Duration
	signingKey string
}

func NewJWTTokenManager(ttl time.Duration, signingKey string) *JWTTokenManager {
	return &JWTTokenManager{
		ttl:        ttl,
		signingKey: signingKey,
	}
}

func (m *JWTTokenManager) GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(m.ttl).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *JWTTokenManager) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
