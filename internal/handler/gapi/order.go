package gapi

import (
	"context"
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type orderHandleGrpc struct {
	pb.UnimplementedOrderServiceServer
	order service.OrderService
}

func NewOrderHandleGrpc(order service.OrderService) *orderHandleGrpc {
	return &orderHandleGrpc{
		order: order,
	}
}

func (h *orderHandleGrpc) GetOrders(ctx context.Context, empty *emptypb.Empty) (*pb.OrderResponses, error) {
	res, err := h.order.GetAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	var orders []*pb.OrderResponse

	for _, o := range *res {
		orders = append(orders, &pb.OrderResponse{
			Id:             int64(o.ID),
			UserId:         int32(o.UserID),
			Name:           o.Name,
			Phone:          o.Phone,
			Email:          o.Email,
			Courier:        o.Courier,
			ShippingMethod: o.ShippingMethod,
			ShippingCost:   int32(o.ShippingCost),
			TotalProduct:   o.TotalProduct,
			TotalPrice:     int32(o.TotalPrice),
			TransactionId:  o.TransactionID,
		})
	}

	return &pb.OrderResponses{
		Orders: orders,
	}, nil
}

func (h *orderHandleGrpc) GetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderRelationResponse, error) {
	res, err := h.order.GetByID(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	var orderItems []*pb.OrderItemResponse

	for _, item := range res.OrderItems {
		orderItems = append(orderItems, &pb.OrderItemResponse{
			Id:       int64(item.ID),
			Name:     item.Name,
			Quantity: int32(item.Quantity),
			Price:    int32(item.Price),
		})
	}

	shippingAddress := pb.ShippingAddressResponse{
		Id:       int64(res.ShippingAddress.ID),
		Alamat:   res.ShippingAddress.Alamat,
		Kota:     res.ShippingAddress.Kota,
		Negara:   res.ShippingAddress.Negara,
		Provinsi: res.ShippingAddress.Provinsi,
	}

	return &pb.OrderRelationResponse{
		Id:              int64(res.ID),
		UserId:          int32(res.UserID),
		Name:            res.Name,
		Phone:           res.Phone,
		Email:           res.Email,
		Courier:         res.Courier,
		ShippingMethod:  res.ShippingMethod,
		ShippingCost:    int32(res.ShippingCost),
		TotalProduct:    res.TotalProduct,
		TotalPrice:      int32(res.TotalPrice),
		TransactionId:   res.TransactionID,
		OrderItems:      orderItems,
		ShippingAddress: &shippingAddress,
	}, nil

}

func (h *orderHandleGrpc) GetOrderUsers(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponses, error) {
	res, err := h.order.OrdersByUser(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	var orders []*pb.OrderResponse

	for _, o := range *res {
		orders = append(orders, &pb.OrderResponse{
			Id:             int64(o.ID),
			UserId:         int32(o.UserID),
			Name:           o.Name,
			Phone:          o.Phone,
			Email:          o.Email,
			Courier:        o.Courier,
			ShippingMethod: o.ShippingMethod,
			ShippingCost:   int32(o.ShippingCost),
			TotalProduct:   o.TotalProduct,
			TotalPrice:     int32(o.TotalPrice),
			TransactionId:  o.TransactionID,
		})
	}

	return &pb.OrderResponses{Orders: orders}, nil
}

func (h *orderHandleGrpc) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {

	createReq := order.CreateOrderRequest{
		UserID:         uint(req.UserId),
		Name:           req.Name,
		Phone:          req.Phone,
		Courier:        req.Courier,
		ShippingMethod: req.ShippingMethod,
		ShippingCost:   int(req.ShippingCost),
		TotalProduct:   req.TotalProduct,
		TotalPrice:     req.TotalPrice,
	}

	shippingReq := order.ShippingAddressRequest{
		Alamat:   req.ShippingAddress.Alamat,
		Kota:     req.ShippingAddress.Kota,
		Negara:   req.ShippingAddress.Negara,
		Provinsi: req.ShippingAddress.Provinsi,
	}

	for _, item := range req.CartItems {
		createReq.CartItems = append(createReq.CartItems, order.CartItemRequest{
			Name:     item.Name,
			Quantity: int(item.Quantity),
			Price:    int(item.Price),
		})
	}

	createReq.ShippingAddress = shippingReq

	if err := createReq.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.order.CreateOrder(&createReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.OrderResponse{
		Id:             int64(res.ID),
		UserId:         int32(res.UserID),
		Name:           res.Name,
		Phone:          res.Phone,
		Email:          res.Email,
		Courier:        res.Courier,
		ShippingMethod: res.ShippingMethod,
		ShippingCost:   int32(res.ShippingCost),
		TotalProduct:   res.TotalProduct,
		TotalPrice:     int32(res.TotalPrice),
		TransactionId:  res.TransactionID,
	}, nil
}
