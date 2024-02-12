package gapi

import (
	"context"
	midtransrequest "ecommerce_fiber/internal/domain/requests/midtrans_request"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type midtransHandleGrpc struct {
	pb.UnimplementedMidtransServiceServer
	midtrans service.MidtransService
}

func NewMidtransHandleGrpc(midtrans service.MidtransService) *midtransHandleGrpc {
	return &midtransHandleGrpc{
		midtrans: midtrans,
	}
}

func (h *midtransHandleGrpc) CreateTransaction(ctx context.Context, req *pb.CreateMidtransRequest) (*pb.SnapResponse, error) {
	createReq := midtransrequest.CreateMidtransRequest{
		GrossAmount: int(req.GrossAmount),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
	}

	if err := createReq.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.midtrans.CreateTransaction(&createReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.SnapResponse{
		Token:         res.Token,
		RedirectUrl:   res.RedirectURL,
		StatusCode:    res.StatusCode,
		ErrorMessages: res.ErrorMessages,
	}, nil
}
