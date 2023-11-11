package app

import (
	"ecommerce_fiber/internal/handler"
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/internal/service"
	"ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/cloudinary"
	"ecommerce_fiber/pkg/database/postgres"
	"ecommerce_fiber/pkg/dotenv"
	"ecommerce_fiber/pkg/hashing"
	"ecommerce_fiber/pkg/logger"

	"github.com/spf13/viper"
)

func Run() {
	log := logger.New(true)

	err := dotenv.Viper()

	if err != nil {
		log.Err(err)
	}

	db, err := postgres.NewClient()

	if err != nil {
		log.Err(err)
	}

	mycloudinary, err := cloudinary.NewMyCloudinary()

	if err != nil {
		log.Err(err)
	}

	hashing := hashing.NewHashingPassword()
	repository := repository.NewRepository(db)
	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		log.Err(err)
	}

	service := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hashing,
		Token:      token,
		Logger:     *log,
	})

	myhandler := handler.NewHandler(service, *mycloudinary, token)

	myhandler.Init().Listen(":8000")
}
