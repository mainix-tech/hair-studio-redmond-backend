package main

import (
	"context"
	"hair-studio-redmond/services/api-gateway/grpc_clients"
	"hair-studio-redmond/shared/env"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	httpAddr = env.GetString("GATEWAY_HTTP_ADDR", ":8081")
)

func mockHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("mockHandler called")
}

type ApiGatewayHandlers struct {
	profileClient *grpc_clients.ProfileServiceClient
	//menuClient    *grpc_clients.MenuServiceClient
	//catalogClient *grpc_clients.CatalogServiceClient
}

func main() {
	log.Printf("Starting API Gateway on %v", httpAddr)

	profileConn, _ := grpc_clients.NewProfileServiceClient()
	defer profileConn.Close()

	handlers := &ApiGatewayHandlers{
		profileClient: profileConn,
		//menuClient:    menuConn,
		//catalogClient: catalogConn,
	}

	mux := http.NewServeMux()

	// Info About the studio domain routes
	mux.HandleFunc("GET /api/v1/profile", mockHandler)
	mux.HandleFunc("PUT /api/v1/profile", handlers.handleUpdateProfileInfo)
	mux.HandleFunc("DELETE /api/v1/profile", mockHandler)
	mux.HandleFunc("POST /api/v1/profile", mockHandler)

	// Services domain routes
	mux.HandleFunc("GET /api/v1/menu", mockHandler)
	mux.HandleFunc("POST /api/v1/menu", mockHandler)
	mux.HandleFunc("PUT /api/v1/menu", mockHandler)
	mux.HandleFunc("DELETE /api/v1/menu", mockHandler)

	// Catalog domain routes
	mux.HandleFunc("GET /api/v1/catalog", mockHandler)
	mux.HandleFunc("POST /api/v1/catalog", mockHandler)
	mux.HandleFunc("PUT /api/v1/catalog", mockHandler)
	mux.HandleFunc("DELETE /api/v1/catalog", mockHandler)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server Listening on %v", httpAddr)
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Printf("Error starting the server: %v", err)

	case sig := <-shutdown:
		log.Printf("Server is shutting down due to %v signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not stop the server gracefully: %v", err)
			server.Close()
		}
	}
}
