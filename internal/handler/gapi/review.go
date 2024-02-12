package gapi

import (
	"context"
	"ecommerce_fiber/internal/domain/requests/review"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type reviewHandleGrpc struct {
	pb.UnimplementedReviewServiceServer
	review service.ReviewService
}

func NewReviewHandleGrpc(review service.ReviewService) *reviewHandleGrpc {
	return &reviewHandleGrpc{
		review: review,
	}
}

func (h *reviewHandleGrpc) GetReviews(ctx context.Context, req *emptypb.Empty) (*pb.ReviewsResponse, error) {
	res, err := h.review.GetAllReviews()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	var reviews []*pb.Review

	for _, review := range *res {
		reviews = append(reviews, &pb.Review{
			Id:      int32(review.ID),
			Name:    review.Name,
			Comment: review.Comment,
			Rating:  int32(review.Rating),
			User: &pb.User{
				Id:        int64(review.User.ID),
				Name:      review.User.Name,
				Email:     review.User.Email,
				CreatedAt: review.User.Created_at,
				UpdatedAt: review.User.Updated_at,
			},
			ProductId: int32(review.ProductID),
			Sentiment: review.Sentiment,
		})
	}

	return &pb.ReviewsResponse{
		Reviews: reviews,
	}, nil

}

func (h *reviewHandleGrpc) GetReview(ctx context.Context, req *pb.ReviewRequest) (*pb.ReviewResponse, error) {
	res, err := h.review.GetReviewByID(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.ReviewResponse{
		Review: &pb.Review{
			Id:      int32(res.ID),
			Name:    res.Name,
			Comment: res.Comment,
			Rating:  int32(res.Rating),
			User: &pb.User{
				Id:        int64(res.User.ID),
				Name:      res.User.Name,
				Email:     res.User.Email,
				CreatedAt: res.User.Created_at,
				UpdatedAt: res.User.Updated_at,
			},
			ProductId: int32(res.ProductID),
			Sentiment: res.Sentiment,
		},
	}, nil

}

func (h *reviewHandleGrpc) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.ReviewResponse, error) {
	createReq := &review.CreateReviewRequest{
		UserID:    int(req.UserId),
		ProductID: int(req.ProductId),
		Rating:    int(req.Rating),
		Comment:   req.Comment,
	}

	if err := createReq.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.review.CreateReview(createReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.ReviewResponse{
		Review: &pb.Review{
			Id:      int32(res.ID),
			Name:    res.Name,
			Comment: res.Comment,
			Rating:  int32(res.Rating),
			User: &pb.User{
				Id:        int64(res.User.ID),
				Name:      res.User.Name,
				Email:     res.User.Email,
				CreatedAt: res.User.Created_at,
				UpdatedAt: res.User.Updated_at,
			},
			ProductId: int32(res.ProductID),
			Sentiment: res.Sentiment,
		},
	}, nil
}
