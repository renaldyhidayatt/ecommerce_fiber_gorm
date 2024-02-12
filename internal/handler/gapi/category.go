package gapi

import (
	"context"
	"ecommerce_fiber/internal/domain/requests/category"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryHandleGrpc struct {
	pb.UnimplementedCategoryServiceServer
	category service.CategoryService
}

func NewCategoryHandleGrpc(category service.CategoryService) *categoryHandleGrpc {
	return &categoryHandleGrpc{
		category: category,
	}
}

func (h *categoryHandleGrpc) GetCategories(ctx context.Context, req *emptypb.Empty) (*pb.CategoriesResponse, error) {

	res, err := h.category.GetCategories()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	var categories []*pb.Category

	for _, category := range res {
		categories = append(categories, &pb.Category{
			Id:        int64(category.ID),
			Name:      category.Name,
			ImagePath: category.ImagePath,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	return &pb.CategoriesResponse{
		Categories: categories,
	}, nil
}

func (h *categoryHandleGrpc) GetCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	res, err := h.category.GetCategory(int(req.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:        int64(res.ID),
			Name:      res.Name,
			ImagePath: res.ImagePath,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	}, nil
}

func (h *categoryHandleGrpc) GetCategorySlug(ctx context.Context, req *pb.CategorySlugRequest) (*pb.CategoryResponse, error) {
	res, err := h.category.GetCategorySlug(req.GetSlug())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:        int64(res.ID),
			Name:      res.Name,
			ImagePath: res.ImagePath,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	}, nil
}

func (h *categoryHandleGrpc) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	newCategory := &category.CreateCategoryRequest{
		Name:     req.Name,
		FilePath: req.FilePath,
	}

	if err := newCategory.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.category.CreateCategory(newCategory)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:        int64(res.ID),
			Name:      res.Name,
			ImagePath: res.ImagePath,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	}, nil
}

func (h *categoryHandleGrpc) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.CategoryResponse, error) {
	reqCategory := &category.UpdateCategoryRequest{
		ID:       int(req.Id),
		Name:     req.Name,
		FilePath: req.FilePath,
	}

	if err := reqCategory.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.category.UpdateCategory(reqCategory)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:        int64(res.ID),
			Name:      res.Name,
			ImagePath: res.ImagePath,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	}, nil
}

func (h *categoryHandleGrpc) DeleteCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.DeleteCategoryResponse, error) {
	_, err := h.category.DeleteCategory(int(req.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.DeleteCategoryResponse{
		Success: true,
	}, nil
}
