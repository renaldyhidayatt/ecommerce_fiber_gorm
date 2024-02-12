package server

import (
	"ecommerce_fiber/internal/handler/gapi"
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/internal/service"
	"ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/database/postgres"
	"ecommerce_fiber/pkg/dotenv"
	"ecommerce_fiber/pkg/hashing"
	"ecommerce_fiber/pkg/logger"
	"flag"
	"fmt"
	"net"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func RunServer() {
	flag.Parse()

	logger, err := logger.NewLogger()

	if err != nil {
		logger.Fatal("Failed to create logger", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	err = dotenv.Viper()

	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
	}

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		logger.Fatal("Failed to create token manager", zap.Error(err))
	}

	conn, err := postgres.NewClient()

	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	hash := hashing.NewHashingPassword()

	repository := repository.NewRepository(conn)

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hash,
		Logger:     *logger,
		Token:      token,
		Mapper:     *mapper,
	})

	grpcServer := grpc.NewServer()

	handleAuth := gapi.NewAuthHandleGrpc(service.Auth)
	handleUser := gapi.NewUserHandleGrpc(service.User)

	pb.RegisterAuthServiceServer(grpcServer, handleAuth)
	pb.RegisterUserServiceServer(grpcServer, handleUser)

	logger.Info(fmt.Sprintf("Server running on port %d", *port))

	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("Failed to serve", zap.Error(err))
	}
}
