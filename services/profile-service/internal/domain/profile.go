package domain

import (
	"context"
	pb "hair-studio-redmond/shared/proto/profile"
)

type ProfileModel struct {
	ID              string `bson:"_id"`
	ProfileEmail    string `bson:"profileEmail"`
	ProfilePhone    string `bson:"profilePhone"`
	ProfileAddress  string `bson:"profileAddress"`
	ProfileTitle    string `bson:"profileTitle"`
	ProfileSubtitle string `bson:"profileSubtitle"`
}

func (t *ProfileModel) ToProto() *pb.ProfileInfo {
	return &pb.ProfileInfo{
		Id:              t.ID,
		ProfileEmail:    t.ProfileEmail,
		ProfilePhone:    t.ProfilePhone,
		ProfileAddress:  t.ProfileAddress,
		ProfileTitle:    t.ProfileTitle,
		ProfileSubtitle: t.ProfileSubtitle,
	}
}

type ProfileRepository interface {
	UpdateProfile(ctx context.Context, dto *ProfileModel) error
}

type ProfileService interface {
	UpdateProfile(ctx context.Context, dto *ProfileModel) error
}
