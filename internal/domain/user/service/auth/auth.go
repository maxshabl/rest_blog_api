package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager interface {
	NewJWT(id int, ttl time.Duration) (string, error)
	ParseJWT(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type AuthManager struct {
	secretKey string
}

func NewAuthManager(secretKey string) (*AuthManager, error) {
	if secretKey == "" {
		return nil, errors.New("empty secret key")
	}
	return &AuthManager{secretKey: secretKey}, nil
}

func (auth AuthManager) NewJWT(id string, ttl time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * ttl)),
		Subject:   id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(auth.secretKey))
}

func (auth AuthManager) ParseJWT(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims["sub"].(string), nil

}

func (auth AuthManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
