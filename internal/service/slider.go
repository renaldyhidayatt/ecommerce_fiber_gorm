package service

import (
	"ecommerce_fiber/internal/domain/requests/slider"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/pkg/logger"

	"go.uber.org/zap"
)

type sliderService struct {
	repository repository.SliderRepository
	logger     logger.Logger
}

func NewSliderService(repository repository.SliderRepository, logger logger.Logger) *sliderService {
	return &sliderService{
		repository: repository,
		logger:     logger,
	}
}

func (s *sliderService) GetAllSliders() (*[]models.Slider, error) {
	res, err := s.repository.GetAllSliders()

	if err != nil {
		s.logger.Error("Error while getting all sliders", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *sliderService) GetSliderByID(sliderID int) (*models.Slider, error) {

	res, err := s.repository.GetSliderByID(sliderID)

	if err != nil {
		s.logger.Error("Error while getting slider by id:", zap.Error(err))
		return nil, err
	}

	return res, nil

}

func (s *sliderService) CreateSlider(request slider.CreateSliderRequest) (*models.Slider, error) {
	schema := &slider.CreateSliderRequest{
		Nama:     request.Nama,
		FilePath: request.FilePath,
	}
	res, err := s.repository.CreateSlider(schema)

	if err != nil {
		s.logger.Error("Error while creating slider:", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *sliderService) UpdateSliderByID(sliderID int, request slider.UpdateSliderRequest) (*models.Slider, error) {
	schema := &slider.UpdateSliderRequest{
		Nama:     request.Nama,
		FilePath: request.FilePath,
	}

	res, err := s.repository.UpdateSliderByID(sliderID, schema)

	if err != nil {
		s.logger.Error("Error while updating slider:", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *sliderService) DeleteSliderByID(sliderID int) (*models.Slider, error) {
	res, err := s.repository.DeleteSliderByID(sliderID)

	if err != nil {
		s.logger.Error("Error while deleting slider:", zap.Error(err))
		return nil, err
	}

	return res, nil
}
