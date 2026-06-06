package domain

import (
	"context"
	pb "hair-studio-redmond/shared/proto/profile"
)

type ProfileModel struct {
	ID          string             `bson:"_id,omitempty"`
	HomePage    HomePageContent    `bson:"homePage"`
	ContactPage ContactPageContent `bson:"contactPage"`
	AboutPage   AboutPageContent   `bson:"aboutPage"`
}

func (t *ProfileModel) ToProto() *pb.ProfileInfo {
	if t == nil {
		return nil
	}

	return &pb.ProfileInfo{
		HomePage: &pb.HomePage{
			Profile:             t.HomePage.Profile,
			ProfileSubtitle:     t.HomePage.ProfileSubtitle,
			AboutStudioTitle:    t.HomePage.AboutTitle,
			AboutStudioSubtitle: t.HomePage.AboutSubtitle,
		},
		ContactPage: &pb.ContactPage{
			ProfileEmail:   t.ContactPage.ProfileEmail,
			ProfilePhone:   t.ContactPage.ProfilePhone,
			ProfileAddress: t.ContactPage.ProfileAddress,
			Title:          t.ContactPage.Title,
			Subtitle:       t.ContactPage.Subtitle,
			WorkHours:      mapWorkHoursToProto(t.ContactPage.WorkHours),
		},
		AboutPage: &pb.AboutPage{
			Title:    t.AboutPage.Title,
			Subtitle: t.AboutPage.Subtitle,
			OurStory: &pb.Story{
				Title:    t.AboutPage.OurStory.Title,
				Subtitle: t.AboutPage.OurStory.Subtitle,
			},
			AboutFounder: &pb.Founder{
				Title:    t.AboutPage.AboutFounder.Title,
				Subtitle: t.AboutPage.AboutFounder.Subtitle,
				Replica:  t.AboutPage.AboutFounder.Replica,
			},
		},
	}
}

// Helper to handle the nested work hours mapping
func mapWorkHoursToProto(wh WorkHours) *pb.WorkHours {
	return &pb.WorkHours{
		Monday:    mapTimeRange(wh.Monday),
		Tuesday:   mapTimeRange(wh.Tuesday),
		Wednesday: mapTimeRange(wh.Wednesday),
		Thursday:  mapTimeRange(wh.Thursday),
		Friday:    mapTimeRange(wh.Friday),
		Saturday:  mapTimeRange(wh.Saturday),
		Sunday:    mapTimeRange(wh.Sunday),
	}
}

func mapTimeRange(tr *TimeRange) *pb.TimeRange {
	if tr == nil {
		return nil
	}
	return &pb.TimeRange{
		Open:  int32(tr.Open),
		Close: int32(tr.Close),
	}
}

type AboutPageContent struct {
	Title        string  `bson:"title"`
	Subtitle     string  `bson:"subtitle"`
	OurStory     Story   `bson:"ourStory"`
	AboutFounder Founder `bson:"aboutFounder"`
}
type Founder struct {
	Title    string `bson:"title"`
	Subtitle string `bson:"subtitle"`
	Replica  string `bson:"replica"`
}
type Story struct {
	Title    string `bson:"title"`
	Subtitle string `bson:"subtitle"`
}
type ContactPageContent struct {
	ProfileEmail   string    `bson:"profileEmail"`
	ProfilePhone   string    `bson:"profilePhone"`
	ProfileAddress string    `bson:"profileAddress"`
	Title          string    `bson:"title"`
	Subtitle       string    `bson:"subtitle"`
	WorkHours      WorkHours `bson:"workHours"`
}
type HomePageContent struct {
	Profile         string `bson:"profile"`
	ProfileSubtitle string `bson:"profileSubtitle"`
	AboutTitle      string `bson:"aboutStudioTitle"`
	AboutSubtitle   string `bson:"aboutStudioSubtitle"`
}
type WorkHours struct {
	Monday    *TimeRange `bson:"monday"`
	Tuesday   *TimeRange `bson:"tuesday"`
	Wednesday *TimeRange `bson:"wednesday"`
	Thursday  *TimeRange `bson:"thoursday"` // Keeping your JSON key typo
	Friday    *TimeRange `bson:"friday"`
	Saturday  *TimeRange `bson:"saturday"`
	Sunday    *TimeRange `bson:"sunday"`
}
type TimeRange struct {
	Open  int `bson:"open"`
	Close int `bson:"close"`
}

type ProfileRepository interface {
	GetProfile(ctx context.Context) (*ProfileModel, error)
	UpdateProfile(ctx context.Context, dto *ProfileModel) error
}

type ProfileService interface {
	GetProfile(ctx context.Context) (*ProfileModel, error)
	UpdateProfile(ctx context.Context, dto *ProfileModel) error
}
