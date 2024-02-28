package api

import (
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/cloudinary"
	mylogger "ecommerce_fiber/pkg/logger"
 _ "ecommerce_fiber/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
"github.com/gofiber/swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type HandlerApi struct {
	tokenManager auth.TokenManager
	cloudinary   cloudinary.MyCloudinary
	auth         pb.AuthServiceClient
	user         pb.UserServiceClient
	category     pb.CategoryServiceClient
	product      pb.ProductServiceClient
	midtrans     pb.MidtransServiceClient
	rajaongkir   pb.RajaOngkirServiceClient
	review       pb.ReviewServiceClient
	slider       pb.SliderServiceClient
}

func NewHandlerApi(conn *grpc.ClientConn, token auth.TokenManager, logger *mylogger.Logger) *HandlerApi {
	clientAuth := pb.NewAuthServiceClient(conn)
	clientUser := pb.NewUserServiceClient(conn)
	clientCategory := pb.NewCategoryServiceClient(conn)
	clientProduct := pb.NewProductServiceClient(conn)
	clientMidtrans := pb.NewMidtransServiceClient(conn)
	clientRajaongkir := pb.NewRajaOngkirServiceClient(conn)
	clientReview := pb.NewReviewServiceClient(conn)
	clientSlider := pb.NewSliderServiceClient(conn)

	mycloudinary, err := cloudinary.NewMyCloudinary()

	if err != nil {
		logger.Fatal("Error while connecting to cloudinary", zap.Error(err))
	}

	return &HandlerApi{
		tokenManager: token,
		auth:         clientAuth,
		user:         clientUser,
		category:     clientCategory,
		product:      clientProduct,
		cloudinary:   *mycloudinary,
		midtrans:     clientMidtrans,
		rajaongkir:   clientRajaongkir,
		review:       clientReview,
		slider:       clientSlider,
	}
}

func (h *HandlerApi) Init() *fiber.App {
	router := fiber.New()

	router.Use(logger.New(
		logger.Config{
			Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "UTC",
		},
	))
	router.Use(cors.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

  router.Get("/docs", swagger.HandlerDefault)

	h.initApi(router)

	return router
}

func (h *HandlerApi) initApi(router *fiber.App) {
	NewAuthHandleApi(h.auth, router)
	NewUserHandleApi(h.user, router)
	NewCategoryHandleApi(h.category, h.cloudinary, router)
	NewProductHandleApi(h.product, h.cloudinary, router)
	NewMidtransHandleApi(h.midtrans, router)
	NewRajaOngkirHandleApi(h.rajaongkir, router)
	NewReviewHandleApi(h.review, router)
	NewSliderHandleApi(h.slider, h.cloudinary, router)

}
