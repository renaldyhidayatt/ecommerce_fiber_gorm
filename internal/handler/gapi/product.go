package gapi

import (
	"context"
	"ecommerce_fiber/internal/domain/requests/product"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productHandleGrpc struct {
	pb.UnimplementedProductServiceServer
	product service.ProductService
}

func NewProductHandleGrpc(product service.ProductService) *productHandleGrpc {
	return &productHandleGrpc{
		product: product,
	}
}

func (h *productHandleGrpc) GetProducts(ctx context.Context, req *emptypb.Empty) (*pb.ProductsResponse, error) {
	res, err := h.product.GetProducts()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	var products []*pb.Product

	for _, product := range res {
		products = append(products, &pb.Product{
			Id:           int32(product.ID),
			Name:         product.Name,
			CategoryId:   int32(product.CategoryID),
			Description:  product.Description,
			CountInStock: int32(product.CountInStock),
			Brand:        product.Brand,
			Weight:       int32(product.Weight),
			Rating:       int32(product.Rating),
			Slug:         product.Slug,
			ImagePath:    product.ImagePath,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
		})
	}

	return &pb.ProductsResponse{
		Products: products,
	}, nil
}

func (h *productHandleGrpc) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	res, err := h.product.GetProduct(int(req.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:           int32(res.ID),
			Name:         res.Name,
			CategoryId:   int32(res.CategoryID),
			Description:  res.Description,
			CountInStock: int32(res.CountInStock),
			Brand:        res.Brand,
			Weight:       int32(res.Weight),
			Rating:       int32(res.Rating),
			Slug:         res.Slug,
			ImagePath:    res.ImagePath,
			CreatedAt:    res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
		},
	}, nil
}

func (h *productHandleGrpc) GetProductSlug(ctx context.Context, req *pb.ProductSlugRequest) (*pb.ProductResponse, error) {
	res, err := h.product.GetProductSlug(req.GetSlug())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:           int32(res.ID),
			Name:         res.Name,
			CategoryId:   int32(res.CategoryID),
			Description:  res.Description,
			CountInStock: int32(res.CountInStock),
			Brand:        res.Brand,
			Weight:       int32(res.Weight),
			Rating:       int32(res.Rating),
			Slug:         res.Slug,
			ImagePath:    res.ImagePath,
			CreatedAt:    res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
		},
	}, nil
}

func (h *productHandleGrpc) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	productReq := &product.CreateProductRequest{
		Name:         req.Name,
		CategoryID:   req.CategoryId,
		Description:  req.Description,
		Price:        int(req.Price),
		Brand:        req.Brand,
		CountInStock: int(req.CountInStock),
		Weight:       int(req.Weight),
		Rating:       int(req.Rating),
		FilePath:     req.FilePath,
	}

	if err := productReq.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.product.CreateProduct(productReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:           int32(res.ID),
			Name:         res.Name,
			CategoryId:   int32(res.CategoryID),
			Description:  res.Description,
			CountInStock: int32(res.CountInStock),
			Brand:        res.Brand,
			Weight:       int32(res.Weight),
			Rating:       int32(res.Rating),
			Slug:         res.Slug,
			ImagePath:    res.ImagePath,
			CreatedAt:    res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
		},
	}, nil
}

func (h *productHandleGrpc) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	productReq := &product.UpdateProductRequest{
		ID:           int(req.Id),
		Name:         req.Name,
		CategoryID:   req.CategoryId,
		Description:  req.Description,
		Price:        int(req.Price),
		Brand:        req.Brand,
		CountInStock: int(req.CountInStock),
		Weight:       int(req.Weight),
		Rating:       int(req.Rating),
		FilePath:     req.FilePath,
	}

	if err := productReq.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.product.UpdateProduct(productReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:           int32(res.ID),
			Name:         res.Name,
			CategoryId:   int32(res.CategoryID),
			Description:  res.Description,
			CountInStock: int32(res.CountInStock),
			Brand:        res.Brand,
			Weight:       int32(res.Weight),
			Rating:       int32(res.Rating),
			Slug:         res.Slug,
			ImagePath:    res.ImagePath,
			CreatedAt:    res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
		},
	}, nil
}

func (h *productHandleGrpc) DeleteProduct(ctx context.Context, req *pb.ProductRequest) (*pb.DeleteProductResponse, error) {
	_, err := h.product.DeleteProduct(int(req.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.DeleteProductResponse{
		Success: true,
	}, nil
}
