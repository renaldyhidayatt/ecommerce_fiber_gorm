package service

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/response"
	userResponse "ecommerce_fiber/internal/domain/response/user"
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/repository"
	customAuth "ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/hashing"
	"ecommerce_fiber/pkg/logger"
	"errors"
	"fmt"
	"strconv"
)

type authService struct {
	user    repository.UserRepository
	hash    hashing.Hashing
	mapping mapper.UserMapping
	log     logger.Logger
	token   customAuth.TokenManager
}

func NewAuthService(user repository.UserRepository, hash hashing.Hashing, log logger.Logger, token customAuth.TokenManager, mapping mapper.UserMapping) *authService {
	return &authService{
		user:    user,
		hash:    hash,
		log:     log,
		token:   token,
		mapping: mapping,
	}
}

func (s *authService) Register(input *auth.RegisterRequest) (*userResponse.UserResponse, error) {
	hashing, err := s.hash.HashPassword(input.Password)

	if err != nil {
		return nil, err
	}

	input.Password = hashing

	user, err := s.user.CreateUser(input)

	if err != nil {
		return nil, err
	}

	return &userResponse.UserResponse{
		ID:         int(user.ID),
		Name:       user.Name,
		Email:      user.Email,
		IsStaff:    user.IsStaff,
		Created_at: user.CreatedAt.String(),
		Updated_at: user.UpdatedAt.String(),
	}, nil
}

func (s *authService) Login(input *auth.LoginRequest) (*response.Token, error) {

	res, err := s.user.GetUserByEmail(input.Email)

	if err != nil {
		return nil, fmt.Errorf("failed error %w", err)
	}

	err = s.hash.ComparePassword(res.Password, input.Password)

	if err != nil {
		return nil, fmt.Errorf("failed error :%w", err)
	}

	if err != nil {
		return nil, err
	}

	access_token, err := s.createAccessToken(int(res.ID), res.Name, res.Name, res.Email, "800000")

	if err != nil {
		return nil, err
	}

	refresh_token, err := s.createRefreshToken(int(res.ID), res.Name, res.Name, res.Email, "800000")

	if err != nil {
		return nil, err
	}

	return &response.Token{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil
}

func (s *authService) RefreshToken(req auth.RefreshTokenRequest) (*response.Token, error) {
	res, err := s.token.ValidateToken(req.RefreshToken)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	idInt, err := strconv.Atoi(res)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	email, err := s.user.GetUser(idInt)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	newToken, err := s.createAccessToken(int(email.ID), email.Name, email.Name, email.Email, "800000")

	if err != nil {
		return nil, errors.New("invalid access token")
	}

	newRefreshToken, err := s.createRefreshToken(int(email.ID), email.Name, email.Name, email.Email, "800000")

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	return &response.Token{
		AccessToken:  newToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *authService) createAccessToken(id int, firstName string, lastName string, email string, phone string) (string, error) {
	res, err := s.token.NewJwtToken(id, firstName, lastName, email, phone, "access")

	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *authService) createRefreshToken(id int, firstName string, lastName string, email string, phone string) (string, error) {
	res, err := s.token.NewJwtToken(id, firstName, lastName, email, phone, "refresh")

	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	return res, nil
}
