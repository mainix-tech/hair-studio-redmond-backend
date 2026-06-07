package main

import (
	menuPb "hair-studio-redmond/shared/proto/menu"
	profilePb "hair-studio-redmond/shared/proto/profile"

	catalogPb "hair-studio-redmond/shared/proto/catalog"
)

// ///////////////////////////////////////////////////// START PROFILE ///////////////////////////////////////////
type UpdateProfileRequest struct {
	HomePage    HomePageContent `json:"homePage"`
	ContactPage ContactPage     `json:"contactPage"`
	AboutPage   AboutPage       `json:"aboutPage"`
}

type HomePageContent struct {
	Profile         string `json:"profile"`
	ProfileSubtitle string `json:"profileSubtitle"`
	AboutTitle      string `json:"aboutStudioTitle"`
	AboutSubtitle   string `json:"aboutStudioSubtitle"`
}

type ContactPage struct {
	ProfileEmail   string    `json:"profileEmail"`
	ProfilePhone   string    `json:"profilePhone"`
	ProfileAddress string    `json:"profileAddress"`
	Title          string    `json:"title"`
	Subtitle       string    `json:"subtitle"`
	WorkHours      WorkHours `json:"workHours"`
}

type WorkHours struct {
	Monday    *TimeRange `json:"monday"`
	Tuesday   *TimeRange `json:"tuesday"`
	Wednesday *TimeRange `json:"wednesday"`
	Thursday  *TimeRange `json:"thoursday"`
	Friday    *TimeRange `json:"friday"`
	Saturday  *TimeRange `json:"saturday"`
	Sunday    *TimeRange `json:"sunday"`
}

type TimeRange struct {
	Open  int `json:"open"`
	Close int `json:"close"`
}

type AboutPage struct {
	Title        string       `json:"title"`
	Subtitle     string       `json:"subtitle"`
	OurStory     Story        `json:"ourStory"`
	AboutFounder AboutFounder `json:"aboutFounder"`
}

type Story struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type AboutFounder struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Replica  string `json:"replica"`
}

// ToProto maps the nested API Gateway request to the gRPC Protobuf structure
func (c *UpdateProfileRequest) ToProto() *profilePb.UpdateProfileInfoRequest {
	return &profilePb.UpdateProfileInfoRequest{
		ProfileInfo: &profilePb.ProfileInfo{
			HomePage: &profilePb.HomePage{
				Profile:             c.HomePage.Profile,
				ProfileSubtitle:     c.HomePage.ProfileSubtitle,
				AboutStudioTitle:    c.HomePage.AboutTitle,
				AboutStudioSubtitle: c.HomePage.AboutSubtitle,
			},
			ContactPage: &profilePb.ContactPage{
				ProfileEmail:   c.ContactPage.ProfileEmail,
				ProfilePhone:   c.ContactPage.ProfilePhone,
				ProfileAddress: c.ContactPage.ProfileAddress,
				Title:          c.ContactPage.Title,
				Subtitle:       c.ContactPage.Subtitle,
				WorkHours:      mapWorkHoursToProto(c.ContactPage.WorkHours),
			},
			AboutPage: &profilePb.AboutPage{
				Title:    c.AboutPage.Title,
				Subtitle: c.AboutPage.Subtitle,
				OurStory: &profilePb.Story{
					Title:    c.AboutPage.OurStory.Title,
					Subtitle: c.AboutPage.OurStory.Subtitle,
				},
				AboutFounder: &profilePb.Founder{
					Title:    c.AboutPage.AboutFounder.Title,
					Subtitle: c.AboutPage.AboutFounder.Subtitle,
					Replica:  c.AboutPage.AboutFounder.Replica,
				},
			},
		},
	}
}

// Helper functions for mapping nested structures
func mapWorkHoursToProto(wh WorkHours) *profilePb.WorkHours {
	return &profilePb.WorkHours{
		Monday:    mapTimeRange(wh.Monday),
		Tuesday:   mapTimeRange(wh.Tuesday),
		Wednesday: mapTimeRange(wh.Wednesday),
		Thursday:  mapTimeRange(wh.Thursday),
		Friday:    mapTimeRange(wh.Friday),
		Saturday:  mapTimeRange(wh.Saturday),
		Sunday:    mapTimeRange(wh.Sunday),
	}
}

func mapTimeRange(tr *TimeRange) *profilePb.TimeRange {
	if tr == nil {
		return nil
	}
	return &profilePb.TimeRange{
		Open:  int32(tr.Open),
		Close: int32(tr.Close),
	}
}

// ///////////////////////////////////////////////////// END PROFILE /////////////////////////////////////////////
// ///////////////////////////////////////////////////// START  MENU /////////////////////////////////////////////
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

/////////////////////////////////////////////////////// END  MENU /////////////////////////////////////////////

// ///////////////////////////////////////////////////// START  CATALOG /////////////////////////////////////////////
type updateCatalogItemRequest struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	ImgUrl   string `json:"imgUrl"`
}
type createCatalogItemRequest struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	ImgUrl   string `json:"imgUrl"`
}
type deleteCatalogItemRequest struct {
	ID string `json:"id"`
}

func (c *createCatalogItemRequest) toProto() *catalogPb.CreateCatalogItemRequest {
	return &catalogPb.CreateCatalogItemRequest{
		Item: &catalogPb.CatalogItem{
			Title:    c.Title,
			Category: c.Category,
			ImgUrl:   c.ImgUrl,
		},
	}
}
func (c *updateCatalogItemRequest) toProto() *catalogPb.UpdateCatalogItemRequest {
	return &catalogPb.UpdateCatalogItemRequest{
		Item: &catalogPb.CatalogItem{
			Id:       c.ID,
			Title:    c.Title,
			ImgUrl:   c.ImgUrl,
			Category: c.Category,
		},
	}
}
func (c *deleteCatalogItemRequest) toProto() *catalogPb.DeleteCatalogItemRequest {
	return &catalogPb.DeleteCatalogItemRequest{
		Id: c.ID,
	}
}
