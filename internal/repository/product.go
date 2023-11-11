package repository

import (
	"ecommerce_fiber/internal/domain/requests/product"
	"ecommerce_fiber/internal/models"
	"errors"
	"strconv"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts() (*[]models.Product, error) {
	var products []models.Product

	db := r.db.Model(&products)

	checkProduct := db.Debug().Find(&products)

	if checkProduct.RowsAffected > 0 {
		return nil, errors.New("row kosong")
	}

	return &products, nil
}

func (r *productRepository) GetProductBySlug(slug string) (*models.Product, error) {
	var product models.Product

	db := r.db.Model(&product)

	checkProductBySlug := db.Preload("Review").Debug().Where("slug_product = ?", slug).First(&product)

	if checkProductBySlug.RowsAffected > 0 {
		return nil, errors.New("failed get slug")
	}

	return &product, nil
}

func (r *productRepository) CreateProduct(request *product.CreateProductRequest) (*models.Product, error) {
	var product models.Product

	slugProduct := slug.Make(request.Name)

	categoryid, err := strconv.Atoi(request.CategoryID)

	db := r.db.Model(&product)

	if err != nil {
		return nil, errors.New("error convert string to int")
	}

	ratingStr := strconv.Itoa(*request.Rating)
	ratingFloat, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		return nil, errors.New("error convert string to float64")
	}

	product.Name = request.Name
	product.SlugProduct = slugProduct
	product.ImageProduct = request.FilePath
	product.Price = request.Price
	product.Weight = request.Weight
	product.Brand = request.Brand
	product.CategoryID = uint(categoryid)
	product.CountInStock = request.CountInStock
	product.Rating = ratingFloat

	addProduct := db.Debug().Create(&product).Commit()

	if addProduct.RowsAffected < 1 {
		return nil, errors.New("error create product")
	}

	return &product, nil
}

func (r *productRepository) GetProductByID(productID int) (*models.Product, error) {
	var product models.Product

	db := r.db.Model(&product)

	checkProductById := db.Debug().Where("id", productID).First(&product)

	if checkProductById.RowsAffected > 0 {
		return nil, errors.New("failed get id")
	}

	return &product, nil
}

func (r *productRepository) MyUpdateQuantity(productID int, quantity int) (bool, error) {
	dbProduct, err := r.GetProductByID(productID)
	if err != nil {
		return false, err
	}

	if dbProduct != nil {
		dbProduct.CountInStock = quantity
		if err := r.db.Save(dbProduct).Error; err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (r *productRepository) UpdateProduct(productID int, request *product.UpdateProductRequest) (*models.Product, error) {
	var product models.Product

	slugProduct := slug.Make(request.Name)

	dbProduct, err := r.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	categoryid, err := strconv.Atoi(request.CategoryID)

	db := r.db.Model(&product)

	if err != nil {
		return nil, errors.New("error convert string to int")
	}

	ratingStr := strconv.Itoa(request.Rating)

	ratingFloat, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		return nil, errors.New("error convert string to float64")
	}

	if dbProduct != nil {
		dbProduct.Name = request.Name
		dbProduct.SlugProduct = slugProduct
		dbProduct.ImageProduct = request.FilePath
		dbProduct.Description = request.Description
		dbProduct.Price = request.Price
		dbProduct.Brand = request.Brand
		dbProduct.Weight = request.Weight
		dbProduct.CategoryID = uint(categoryid)
		dbProduct.Brand = request.Brand
		dbProduct.CountInStock = request.CountInStock
		dbProduct.Rating = ratingFloat

		if err := db.Debug().Save(dbProduct).Error; err != nil {
			return nil, err
		}

		return dbProduct, nil
	}
	return nil, nil
}

func (r *productRepository) DeleteProduct(productID int) (*models.Product, error) {
	var product models.Product

	dbProduct, err := r.GetProductByID(productID)

	if err != nil {
		return nil, err
	}

	if dbProduct == nil {
		return nil, errors.New("product not found")
	}

	db := r.db.Model(&product)
	if err := db.Debug().Delete(dbProduct).Error; err != nil {
		return nil, err
	}

	return dbProduct, nil
}
