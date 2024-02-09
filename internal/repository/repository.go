package repository

import "gorm.io/gorm"

type Repositories struct {
	Cart     CartRepository
	Category CategoryRepository
	Order    OrderRepository
	Product  ProductRepository
	Review   ReviewRepository
	Slider   SliderRepository
	User     UserRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepository(db),
		Category: NewCategoryRepository(db),
		Product:  NewProductRepository(db),
		Cart:     NewCartRepository(db),
		Order:    NewOrderRepository(db),
		Review:   NewReviewRepository(db),
		Slider:   NewSliderRepository(db),
	}
}
