package gapi

import (
	"context"
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/requests/user"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userHandleGrpc struct {
	pb.UnimplementedUserServiceServer
	user service.UserService
}

func NewUserHandleGrpc(user service.UserService) *userHandleGrpc {
	return &userHandleGrpc{
		user: user,
	}
}

func (h *userHandleGrpc) GetUsers(ctx context.Context, empty *emptypb.Empty) (*pb.UsersResponse, error) {
	res, err := h.user.GetUsers()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	var users []*pb.User

	for _, user := range *res {
		users = append(users, &pb.User{
			Id:        int64(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			IsStaff:   user.IsStaff,
			CreatedAt: user.Created_at,
			UpdatedAt: user.Updated_at,
		})
	}

	return &pb.UsersResponse{
		Users: users,
	}, nil
}

func (h *userHandleGrpc) GetUser(ctx context.Context, id *pb.UserRequest) (*pb.UserResponse, error) {
	res, err := h.user.GetUser(int(id.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:        int64(res.ID),
			Name:      res.Name,
			Email:     res.Email,
			IsStaff:   res.IsStaff,
			CreatedAt: res.Created_at,
			UpdatedAt: res.Updated_at,
		},
	}, nil

}

func (h *userHandleGrpc) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := auth.RegisterRequest{
		Name:             req.Name,
		Email:            req.Email,
		Password:         req.Password,
		Confirm_password: req.ConfirmPassword,
	}

	if err := user.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.user.CreateUser(&user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:        int64(res.ID),
			Name:      res.Name,
			Email:     res.Email,
			IsStaff:   res.IsStaff,
			CreatedAt: res.Created_at,
			UpdatedAt: res.Updated_at,
		},
	}, nil
}

func (h *userHandleGrpc) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	reqUser := user.UpdateUserRequest{
		ID:       int(req.Id),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := reqUser.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.user.UpdateUser(&reqUser)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:        int64(res.ID),
			Name:      res.Name,
			Email:     res.Email,
			IsStaff:   res.IsStaff,
			CreatedAt: res.Created_at,
			UpdatedAt: res.Updated_at,
		},
	}, nil

}

func (h *userHandleGrpc) DeleteUser(ctx context.Context, id *pb.UserRequest) (*pb.DeleteUserResponse, error) {
	_, err := h.user.DeleteUser(int(id.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}
