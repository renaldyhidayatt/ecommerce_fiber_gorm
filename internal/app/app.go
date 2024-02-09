package app

import (
	"ecommerce_fiber/internal/handler"
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/internal/service"
	"ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/cloudinary"
	"ecommerce_fiber/pkg/database/postgres"
	"ecommerce_fiber/pkg/dotenv"
	"ecommerce_fiber/pkg/hashing"
	"ecommerce_fiber/pkg/logger"
	midtranspkg "ecommerce_fiber/pkg/midtrans_pkg"
	"ecommerce_fiber/pkg/rajaongkir"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Run() {
	log, err := logger.NewLogger()

	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	err = dotenv.Viper()

	if err != nil {
		log.Error("Error while loading .env file", zap.Error(err))
	}

	db, err := postgres.NewClient()

	if err != nil {
		log.Error("Error while connecting to database", zap.Error(err))
	}

	mycloudinary, err := cloudinary.NewMyCloudinary()

	if err != nil {
		log.Error("Error while connecting to cloudinary", zap.Error(err))
	}

	hashing := hashing.NewHashingPassword()
	repository := repository.NewRepository(db)
	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		log.Error("Error while initializing token manager", zap.Error(err))
	}

	rajaOngkir := rajaongkir.NewRajaOngkirAPI()

	snapMidtrans := midtranspkg.NewSnapClient()

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hashing,
		Token:      token,
		Logger:     *log,
		Snap:       snapMidtrans,
		RajaOngkir: rajaOngkir,
		Mapper:     *mapper,
	})

	myhandler := handler.NewHandler(service, *mycloudinary, token)

	myhandler.Init().Listen(":8000")
}
