package main

import (
	"context"
	"hair-studio-redmond/services/profile-service/internal/infrastructure/db"
	"hair-studio-redmond/services/profile-service/internal/infrastructure/repository"
	"hair-studio-redmond/services/profile-service/internal/service"
	"log"
	"net"
	"os"
	"os/signal"

	"embed"
	"hair-studio-redmond/services/profile-service/internal/infrastructure/grpc"
	"syscall"

	grpcserver "google.golang.org/grpc"
)

var GrpcAddr = ":9093"

//go:embed migrations
var migrationFiles embed.FS

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Establish the database connection pool
	dbConn, err := db.ConnectPostgres(ctx, migrationFiles)
	if err != nil {
		log.Fatalf("Critical database error: %v", err)
	}

	defer dbConn.Close(ctx) // Ensure resources clear out gracefully on shutdown
	repo := repository.NewPostgresRepository(dbConn)
	svc := service.NewService(repo)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Starting the gRPC server
	grpcServer := grpcserver.NewServer()
	grpc.NewGRPCHandler(grpcServer, svc)

	log.Printf("Starting gRPC server Trip service on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	// wait for the shutdown signal
	<-ctx.Done()
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}
