package client

import (
	"context"
	"ecommerce_fiber/internal/handler/api"
	"ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/dotenv"
	"ecommerce_fiber/pkg/logger"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

// @title EcommerceFiber
// @version 1.0
// @description REST API for Ecommerce fiber

// @host localhost:8000
// @BasePath /api/

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization

// Run initializes whole application.
func RunClient() {
	flag.Parse()

	logger, err := logger.NewLogger()

	if err != nil {
		logger.Fatal("Failed to create logger", zap.Error(err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("Failed to connect to server", zap.Error(err))
	}

	if err != nil {
		logger.Fatal("Failed to connect to server", zap.Error(err))
	}

	err = dotenv.Viper()

	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
	}

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		logger.Fatal("Failed to create token manager", zap.Error(err))
	}

	handler := api.NewHandlerApi(conn, token, logger)

	go func() {
		if err := handler.Init().Listen(":8000"); err != nil {
			logger.Error("server error", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down the server...")

	cancel()

	logger.Info("Server has been shut down gracefully")
}
