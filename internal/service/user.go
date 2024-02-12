package service

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/requests/user"
	userResponse "ecommerce_fiber/internal/domain/response/user"
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/repository"
	"errors"

	"ecommerce_fiber/pkg/hashing"
	"ecommerce_fiber/pkg/logger"
)

type userService struct {
	Repository repository.UserRepository
	mapper     mapper.UserMapping
	hash       hashing.Hashing
	log        logger.Logger
}

func NewUserService(auth repository.UserRepository, hash hashing.Hashing, logger logger.Logger, mapper mapper.UserMapping) *userService {
	return &userService{
		Repository: auth,
		hash:       hash,
		log:        logger,
		mapper:     mapper,
	}
}

func (s *userService) GetUsers() (*[]userResponse.UserResponse, error) {

	res, err := s.Repository.GetUsers()

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponses(res)

	return mapper, nil
}

func (s *userService) CreateUser(request *auth.RegisterRequest) (*userResponse.UserResponse, error) {

	hashing, err := s.hash.HashPassword(request.Password)

	if err != nil {
		return nil, errors.New("error hashing password")
	}

	request.Password = hashing

	res, err := s.Repository.CreateUser(request)

	if err != nil {
		return nil, errors.New("error creating user")
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) GetUser(id int) (*userResponse.UserResponse, error) {
	res, err := s.Repository.GetUser(id)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) UpdateUser(request *user.UpdateUserRequest) (*userResponse.UserResponse, error) {
	hashing, err := s.hash.HashPassword(request.Password)

	if err != nil {
		return nil, err
	}

	request.Password = hashing

	res, err := s.Repository.UpdateUser(request)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) DeleteUser(id int) (*userResponse.UserResponse, error) {
	res, err := s.Repository.DeleteUser(id)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}
