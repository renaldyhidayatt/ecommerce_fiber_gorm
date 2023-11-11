package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	NewJWT(userId int) (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(userId int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expireTime.Unix(),
		"sub": fmt.Sprintf("%d", userId),
	})

	return token.SignedString([]byte(m.signingKey))
}
