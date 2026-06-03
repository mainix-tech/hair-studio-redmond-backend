package main

import pb "hair-studio-redmond/shared/proto/profile"

type updateProfileInfoRequest struct {
	ID              string `json:"id"`
	ProfileEmail    string `json:"profileEmail"`
	ProfilePhone    string `json:"profilePhone"`
	ProfileAddress  string `json:"profileAddress"`
	ProfileTitle    string `json:"profile"`
	ProfileSubtitle string `json:"profileSubtitle"`
}

func (c *updateProfileInfoRequest) toProto() *pb.UpdateProfileInfoRequest {
	return &pb.UpdateProfileInfoRequest{
		ProfileInfo: &pb.ProfileInfo{
			Id:              c.ID,
			ProfileEmail:    c.ProfileEmail,
			ProfilePhone:    c.ProfilePhone,
			ProfileAddress:  c.ProfileAddress,
			ProfileTitle:    c.ProfileTitle,
			ProfileSubtitle: c.ProfileSubtitle,
		},
	}
}
