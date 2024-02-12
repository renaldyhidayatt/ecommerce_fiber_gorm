package gapi

import (
	"context"
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authHandleGrpc struct {
	pb.UnimplementedAuthServiceServer
	auth service.AuthService
}

func NewAuthHandleGrpc(auth service.AuthService) *authHandleGrpc {
	return &authHandleGrpc{
		auth: auth,
	}
}

func (a *authHandleGrpc) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := auth.RegisterRequest{
		Name:             req.Name,
		Email:            req.Email,
		Password:         req.Password,
		Confirm_password: req.ConfirmPassword,
	}

	if err := user.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := a.auth.Register(&user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.RegisterResponse{
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

func (a *authHandleGrpc) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	user := auth.LoginRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := user.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := a.auth.Login(&user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}

func (a *authHandleGrpc) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.AuthResponse, error) {
	token := auth.RefreshTokenRequest{
		RefreshToken: req.GetRefreshToken(),
	}

	res, err := a.auth.RefreshToken(token)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
