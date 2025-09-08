package main

import (
	"log"
	"net"

	"go-grpc-crud/internal/config"
	"go-grpc-crud/internal/db"
	"go-grpc-crud/internal/repository"
	"go-grpc-crud/internal/service"
	"go-grpc-crud/proto/userpb"

	"google.golang.org/grpc"
)

func main() {
	// 1. Cargar config
	cfg := config.LoadConfig()

	// 2. Conectar DB con config
	db.ConnectGorm(cfg)

	// 3. Servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userRepo := repository.NewUserRepositoryGorm()
	userService := service.NewUserServiceServer(userRepo)

	userpb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
