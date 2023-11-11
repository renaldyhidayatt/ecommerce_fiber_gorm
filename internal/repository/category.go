package repository

import (
	"ecommerce_fiber/internal/domain/requests/category"
	"ecommerce_fiber/internal/models"
	"errors"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll() (*[]models.Category, error) {
	var categories []models.Category

	db := r.db.Model(&categories)

	checkCategory := db.Debug().Find(&categories)

	if checkCategory.RowsAffected < 1 {
		return nil, errors.New("row kosong")
	}

	return &categories, nil

}

func (r *categoryRepository) GetByID(categoryID int) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(category)

	checkCategoryById := db.Debug().Where("id = ?", categoryID).First(&category)

	if checkCategoryById.RowsAffected > 0 {
		return &category, errors.New("error not found category")
	}

	return &category, nil
}

func (r *categoryRepository) GetBySlug(slug string) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(&category)

	checkCategorySlug := db.Debug().Where("slug_category", slug).First(&category)

	if checkCategorySlug.RowsAffected > 0 {

		return &category, errors.New("error check category already")
	}

	return &category, nil
}

func (r *categoryRepository) Create(request *category.CreateCategoryRequest) (*models.Category, error) {
	var myCategory models.Category

	slugCategory := slug.Make(request.Name)

	myCategory.Name = request.Name
	myCategory.SlugCategory = slugCategory
	myCategory.Description = "hello"
	myCategory.ImageCategory = request.FilePath

	db := r.db.Model(&myCategory)

	checkCategoryName := db.Debug().Where("name = ?", request.Name).First(&myCategory)
	if checkCategoryName.RowsAffected > 0 {
		return &myCategory, errors.New("error check category already exists")
	}

	addCategory := db.Debug().Create(&myCategory).Commit()
	if addCategory.RowsAffected < 1 {
		return &myCategory, errors.New("error creating category")
	}

	return &myCategory, nil
}

func (r *categoryRepository) UpdateByID(id int, updatedCategory *category.UpdateCategoryRequest) (*models.Category, error) {
	var category models.Category

	slugCategory := slug.Make(updatedCategory.Name)

	db := r.db.Model(&category)

	checkCategoryById := db.Debug().Where("id = ?", id).First(&category)

	if checkCategoryById.RowsAffected > 1 {
		return &category, errors.New("error not found category")
	}

	category.Name = updatedCategory.Name
	category.SlugCategory = slugCategory
	category.Description = "hello"
	category.ImageCategory = updatedCategory.FilePath

	updateCategory := db.Debug().Updates(&category)

	if updateCategory.RowsAffected > 1 {
		return &category, errors.New("error Failed update")
	}

	return &category, nil

}

func (r *categoryRepository) DeleteByID(categoryID int) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(&category)

	checkCategoryById := db.Debug().Where("id =?", categoryID).First(&category)

	if checkCategoryById.RowsAffected > 1 {
		return &category, errors.New("error not found category")
	}

	deleteCategory := db.Debug().Delete(&category)

	if deleteCategory.RowsAffected > 1 {
		return &category, errors.New("failed delete category")
	}

	return &category, nil
}
