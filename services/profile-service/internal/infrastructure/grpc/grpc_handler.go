package grpc

import (
	"context"
	"hair-studio-redmond/services/profile-service/internal/domain"
	pb "hair-studio-redmond/shared/proto/profile"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedProfileServiceServer

	service domain.ProfileService
}

func NewGRPCHandler(server *grpc.Server, service domain.ProfileService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}

	pb.RegisterProfileServiceServer(server, handler)

	return handler
}

func (h *gRPCHandler) UpdateProfileInfo(ctx context.Context, req *pb.UpdateProfileInfoRequest) (*pb.UpdateProfileInfoResponse, error) {
	log.Printf("UpdateProfileInfoRequest %+v", req)

	if err := h.service.UpdateProfile(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create the trip: %v", err)
	}

	return &pb.UpdateProfileInfoResponse{}, nil
}
