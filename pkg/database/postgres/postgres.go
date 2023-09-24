package postgres

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), viper.GetString("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed error connect to database")
	}

	// err = db.AutoMigrate(&models.ModelCategory{}, &models.ModelCustomer{}, &models.ModelProduct{}, &models.ModelProductKeluar{}, &models.ModelProductMasuk{}, &models.ModelSale{}, &models.ModelSupplier{}, &models.ModelUser{})

	if err != nil {

		return nil, fmt.Errorf("faield connected to database")
	}

	return db, nil
}
