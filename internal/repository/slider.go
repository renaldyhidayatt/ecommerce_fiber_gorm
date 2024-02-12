package repository

import (
	"ecommerce_fiber/internal/domain/requests/slider"
	"ecommerce_fiber/internal/models"
	"errors"

	"gorm.io/gorm"
)

type sliderRepository struct {
	db *gorm.DB
}

func NewSliderRepository(db *gorm.DB) *sliderRepository {
	return &sliderRepository{db: db}
}

func (r *sliderRepository) GetAllSliders() (*[]models.Slider, error) {
	var sliders []models.Slider

	db := r.db.Model(&sliders)

	checkSlider := db.Debug().Find(&sliders)

	if checkSlider.RowsAffected < 0 {
		return nil, errors.New("row kosong")
	}

	return &sliders, nil
}

func (r *sliderRepository) GetSliderByID(sliderID int) (*models.Slider, error) {
	var slider models.Slider

	db := r.db.Model(slider)

	checkSliderById := db.Debug().Where("id = ?", sliderID).First(&slider)

	if checkSliderById.RowsAffected < 0 {
		return nil, errors.New("failed get id")
	}

	return &slider, nil
}

func (r *sliderRepository) CreateSlider(sliderRequest *slider.CreateSliderRequest) (*models.Slider, error) {
	var slider models.Slider

	db := r.db.Model(&slider)

	mySlider := models.Slider{
		Name:  sliderRequest.Nama,
		Image: sliderRequest.FilePath,
	}

	addSlider := db.Debug().Create(&mySlider).Commit()

	if addSlider.RowsAffected < 1 {
		return &mySlider, errors.New("error create slider")
	}

	return &mySlider, nil
}

func (r *sliderRepository) UpdateSliderByID(updatedSlider *slider.UpdateSliderRequest) (*models.Slider, error) {
	mySlider, err := r.GetSliderByID(updatedSlider.ID)
	if err != nil {
		return nil, err
	}

	if mySlider != nil {
		mySlider.Name = updatedSlider.Nama
		mySlider.Image = updatedSlider.FilePath

		if err := r.db.Debug().Save(mySlider).Error; err != nil {
			return nil, err
		}

		return mySlider, nil
	}
	return nil, errors.New("slider not found")
}

func (r *sliderRepository) DeleteSliderByID(sliderID int) (*models.Slider, error) {
	var slider models.Slider

	mySlider, err := r.GetSliderByID(sliderID)
	if err != nil {
		return nil, err
	}

	if mySlider == nil {
		return nil, errors.New("slider not found")
	}

	db := r.db.Model(&slider)

	if err := db.Debug().Delete(mySlider).Error; err != nil {
		return nil, err
	}

	return mySlider, nil
}
