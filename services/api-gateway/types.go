package main

import profilePb "hair-studio-redmond/shared/proto/profile"
import menuPb "hair-studio-redmond/shared/proto/menu"

type updateProfileInfoRequest struct {
	ID              string `json:"id"`
	ProfileEmail    string `json:"profileEmail"`
	ProfilePhone    string `json:"profilePhone"`
	ProfileAddress  string `json:"profileAddress"`
	ProfileTitle    string `json:"profile"`
	ProfileSubtitle string `json:"profileSubtitle"`
}

func (c *updateProfileInfoRequest) toProto() *profilePb.UpdateProfileInfoRequest {
	return &profilePb.UpdateProfileInfoRequest{
		ProfileInfo: &profilePb.ProfileInfo{
			Id:              c.ID,
			ProfileEmail:    c.ProfileEmail,
			ProfilePhone:    c.ProfilePhone,
			ProfileAddress:  c.ProfileAddress,
			ProfileTitle:    c.ProfileTitle,
			ProfileSubtitle: c.ProfileSubtitle,
		},
	}
}

type updateMenuItemRequest struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	PreviewDescription  string `json:"previewDescription"`
	DetailedDescription string `json:"detailedDescription"`
	Time                int64  `json:"time"`
	Price               int64  `json:"price"`
	Category            string `json:"category"`
}

type createMenuItemRequest struct {
	Name                string `json:"name"`
	PreviewDescription  string `json:"previewDescription"`
	DetailedDescription string `json:"detailedDescription"`
	Time                int64  `json:"time"`
	Price               int64  `json:"price"`
	Category            string `json:"category"`
}
type deleteMenuItemRequest struct {
	ID string `json:"id"`
}

func (c *createMenuItemRequest) toProto() *menuPb.CreateMenuItemRequest {
	return &menuPb.CreateMenuItemRequest{
		Item: &menuPb.MenuItem{
			Name:                c.Name,
			PreviewDescription:  c.PreviewDescription,
			DetailedDescription: c.DetailedDescription,
			Time:                c.Time,
			Price:               c.Price,
			Category:            c.Category,
		},
	}
}
func (c *updateMenuItemRequest) toProto() *menuPb.UpdateMenuItemRequest {
	return &menuPb.UpdateMenuItemRequest{
		Item: &menuPb.MenuItem{
			Id:                  c.ID,
			Name:                c.Name,
			PreviewDescription:  c.PreviewDescription,
			DetailedDescription: c.DetailedDescription,
			Time:                c.Time,
			Price:               c.Price,
			Category:            c.Category,
		},
	}
}
func (c *deleteMenuItemRequest) toProto() *menuPb.DeleteMenuItemRequest {
	return &menuPb.DeleteMenuItemRequest{
		Id: c.ID,
	}
}
