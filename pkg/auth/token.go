package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	jwt.RegisteredClaims
}

type TokenManager interface {
	NewJwtToken(id int, firstName string, lastName string, email string, phone string, audience string) (string, error)
	ValidateToken(accessToken string) (string, error)
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

func (m *Manager) NewJwtToken(id int, firstName string, lastName string, email string, phone string, audience string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtCustomClaims{
		firstName,
		lastName,
		email,
		phone,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    strconv.Itoa(int(nowTime.Unix())),
			Subject:   strconv.Itoa(id),
			Audience:  []string{audience},
		},
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) ValidateToken(accessToken string) (string, error) {
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
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("subject claim not found or not a string")
	}

	fmt.Println("Subject: ", subject)

	return subject, nil
}
