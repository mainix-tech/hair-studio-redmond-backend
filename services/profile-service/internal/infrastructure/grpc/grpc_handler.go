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

	// 1. Safety check: Verify the incoming nested ProfileInfo object is not nil
	info := req.GetProfileInfo()
	if info == nil {
		return nil, status.Error(codes.InvalidArgument, "profile_info payload cannot be empty")
	}

	// 2. Unpack the proto data fields into your reusable Domain model
	model := &domain.ProfileModel{
		ID:              info.GetId(),
		ProfileEmail:    info.GetProfileEmail(),
		ProfilePhone:    info.GetProfilePhone(),
		ProfileAddress:  info.GetProfileAddress(),
		ProfileTitle:    info.GetProfileTitle(),
		ProfileSubtitle: info.GetProfileSubtitle(),
	}

	// 3. Pass the pointer to your domain model into your service layer
	if err := h.service.UpdateProfile(ctx, model); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update profile: %v", err)
	}

	return &pb.UpdateProfileInfoResponse{}, nil
}
