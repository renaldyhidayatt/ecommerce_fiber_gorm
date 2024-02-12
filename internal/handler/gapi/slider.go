package gapi

import (
	"context"
	"ecommerce_fiber/internal/domain/requests/slider"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type sliderHandleGrpc struct {
	pb.UnimplementedSliderServiceServer
	slider service.SliderService
}

func NewSliderHandleGrpc(slider service.SliderService) *sliderHandleGrpc {
	return &sliderHandleGrpc{
		slider: slider,
	}
}

func (h *sliderHandleGrpc) GetSliders(ctx context.Context, req *emptypb.Empty) (*pb.SlidersResponse, error) {
	res, err := h.slider.GetAllSliders()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	var sliders []*pb.Slider

	for _, s := range *res {
		sliders = append(sliders, &pb.Slider{
			Id:    int64(s.ID),
			Name:  s.Name,
			Image: s.Image,
		})
	}

	return &pb.SlidersResponse{
		Sliders: sliders,
	}, nil
}

func (h *sliderHandleGrpc) GetSlider(ctx context.Context, req *pb.SliderRequest) (*pb.SliderResponse, error) {
	res, err := h.slider.GetSliderByID(int(req.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.SliderResponse{
		Slider: &pb.Slider{
			Id:    int64(res.ID),
			Name:  res.Name,
			Image: res.Image,
		},
	}, nil
}

func (h *sliderHandleGrpc) CreateSlider(ctx context.Context, req *pb.CreateSliderRequest) (*pb.SliderResponse, error) {
	createReq := slider.CreateSliderRequest{
		Nama:     req.Name,
		FilePath: req.Image,
	}

	if err := createReq.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %s", err.Error())
	}

	res, err := h.slider.CreateSlider(createReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.SliderResponse{
		Slider: &pb.Slider{
			Id:    int64(res.ID),
			Name:  res.Name,
			Image: res.Image,
		},
	}, nil
}

func (h *sliderHandleGrpc) UpdateSlider(ctx context.Context, req *pb.UpdateSliderRequest) (*pb.SliderResponse, error) {
	updateReq := slider.UpdateSliderRequest{
		ID:       int(req.Id),
		Nama:     req.Name,
		FilePath: req.Image,
	}

	res, err := h.slider.UpdateSliderByID(updateReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.SliderResponse{
		Slider: &pb.Slider{
			Id:    int64(res.ID),
			Name:  res.Name,
			Image: res.Image,
		},
	}, nil
}

func (h *sliderHandleGrpc) DeleteSlider(ctx context.Context, req *pb.SliderRequest) (*pb.DeleteSliderResponse, error) {
	_, err := h.slider.DeleteSliderByID(int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	return &pb.DeleteSliderResponse{
		Success: true,
	}, nil
}
