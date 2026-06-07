package grpc

import (
	"context"
	"hair-studio-redmond/service/profile-service/internal/domain"
	pb "hair-studio-redmond/shared/proto/profile"

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

func (h *gRPCHandler) GetProfileInfo(ctx context.Context, req *pb.GetProfileInfoRequest) (*pb.GetProfileInfoResponse, error) {
	// 1. Call the service layer to get the domain model
	profile, err := h.service.GetProfile(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch profile: %v", err)
	}

	// 2. Check if the profile was found
	if profile == nil {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	// 3. Use your helper method to convert domain model to Protobuf
	return &pb.GetProfileInfoResponse{
		ProfileInfo: profile.ToProto(),
	}, nil
}
func (h *gRPCHandler) UpdateProfileInfo(ctx context.Context, req *pb.UpdateProfileInfoRequest) (*pb.UpdateProfileInfoResponse, error) {
	info := req.GetProfileInfo()
	if info == nil {
		return nil, status.Error(codes.InvalidArgument, "profile_info payload cannot be empty")
	}

	model := &domain.ProfileModel{
		HomePage: domain.HomePageContent{
			Profile:         info.HomePage.GetProfile(),
			ProfileSubtitle: info.HomePage.GetProfileSubtitle(),
			AboutTitle:      info.HomePage.GetAboutStudioTitle(),
			AboutSubtitle:   info.HomePage.GetAboutStudioSubtitle(),
		},
		ContactPage: domain.ContactPageContent{
			ProfileEmail:   info.ContactPage.GetProfileEmail(),
			ProfilePhone:   info.ContactPage.GetProfilePhone(),
			ProfileAddress: info.ContactPage.GetProfileAddress(),
			Title:          info.ContactPage.GetTitle(),
			Subtitle:       info.ContactPage.GetSubtitle(),
			WorkHours:      mapWorkHoursFromProto(info.ContactPage.GetWorkHours()),
		},
		AboutPage: domain.AboutPageContent{
			Title:    info.AboutPage.GetTitle(),
			Subtitle: info.AboutPage.GetSubtitle(),
			OurStory: domain.Story{
				Title:    info.AboutPage.OurStory.GetTitle(),
				Subtitle: info.AboutPage.OurStory.GetSubtitle(),
			},
			AboutFounder: domain.Founder{
				Title:    info.AboutPage.AboutFounder.GetTitle(),
				Subtitle: info.AboutPage.AboutFounder.GetSubtitle(),
				Replica:  info.AboutPage.AboutFounder.GetReplica(),
			},
		},
	}

	// 2. Pass the fully populated domain model to your service layer
	if err := h.service.UpdateProfile(ctx, model); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update profile: %v", err)
	}

	return &pb.UpdateProfileInfoResponse{}, nil
}

func mapWorkHoursFromProto(wh *pb.WorkHours) domain.WorkHours {
	if wh == nil {
		return domain.WorkHours{}
	}
	return domain.WorkHours{
		Monday:    mapTimeRangeFromProto(wh.Monday),
		Tuesday:   mapTimeRangeFromProto(wh.Tuesday),
		Wednesday: mapTimeRangeFromProto(wh.Wednesday),
		Thursday:  mapTimeRangeFromProto(wh.Thursday),
		Friday:    mapTimeRangeFromProto(wh.Friday),
		Saturday:  mapTimeRangeFromProto(wh.Saturday),
		Sunday:    mapTimeRangeFromProto(wh.Sunday),
	}
}
func mapTimeRangeFromProto(tr *pb.TimeRange) *domain.TimeRange {
	if tr == nil {
		return nil
	}
	return &domain.TimeRange{
		Open:  int(tr.Open),
		Close: int(tr.Close),
	}
}
