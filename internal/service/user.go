package service

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/requests/user"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"

	customAuth "ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/hashing"
	"ecommerce_fiber/pkg/logger"
	"fmt"
)

type userService struct {
	Repository repository.UserRepository
	hash       hashing.Hashing
	log        logger.Logger
	token      customAuth.TokenManager
}

func NewUserService(auth repository.UserRepository, hash hashing.Hashing, token customAuth.TokenManager, logger logger.Logger) *userService {
	return &userService{
		Repository: auth,
		hash:       hash,
		token:      token,
		log:        logger,
	}
}

func (s *userService) Register(input *auth.RegisterRequest) (*models.User, error) {
	hashing, err := s.hash.HashPassword(input.Password)

	if err != nil {
		return nil, err
	}

	input.Password = hashing

	user, err := s.Repository.CreateUser(input)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(input *auth.LoginRequest) (auth.Token, error) {

	res, err := s.Repository.GetUserByEmail(input.Email)

	if err != nil {
		return auth.Token{}, fmt.Errorf("failed error %w", err)
	}

	err = s.hash.ComparePassword(res.Password, input.Password)

	if err != nil {
		return auth.Token{}, fmt.Errorf("failed error :%w", err)
	}

	if err != nil {
		return auth.Token{}, err
	}

	return s.createJwt(int(res.ID))
}

func (s *userService) createJwt(id int) (auth.Token, error) {
	var (
		res auth.Token
		err error
	)

	res.Jwt, err = s.token.NewJWT(id)

	if err != nil {
		return res, err
	}

	return res, err
}

func (s *userService) GetUserAll() (*[]models.User, error) {

	res, err := s.Repository.GetUserAll()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) GetUserById(id int) (*models.User, error) {
	res, err := s.Repository.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) UpdateUserById(id int, request *user.UpdateUserRequest) (*models.User, error) {
	hashing, err := s.hash.HashPassword(request.Password)

	if err != nil {
		return nil, err
	}

	request.Password = hashing

	res, err := s.Repository.UpdateUserById(id, request)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) DeleteUserById(id int) (*models.User, error) {
	res, err := s.Repository.DeleteUserById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
