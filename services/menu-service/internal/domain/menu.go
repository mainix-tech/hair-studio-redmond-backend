package domain

import (
	"context"
	pb "hair-studio-redmond/shared/proto/menu"
)

// 1. Core Model: ONLY represents a full data record that exists in the DB (Has an ID!)
type MenuModel struct {
	ID                  string `bson:"_id"`
	Name                string `bson:"name"`
	PreviewDescription  string `bson:"previewDescription"`
	DetailedDescription string `bson:"detailedDescription"`
	Time                int64  `bson:"time"`
	Price               int64  `bson:"price"`
	Category            string `bson:"category"`
}

type MenuCreate struct {
	Name                string
	PreviewDescription  string
	DetailedDescription string
	Time                int64
	Price               int64
	Category            string
}

func (t *MenuModel) ToProto() *pb.MenuItem {
	return &pb.MenuItem{
		Id:                  t.ID,
		Name:                t.Name,
		PreviewDescription:  t.PreviewDescription,
		DetailedDescription: t.DetailedDescription,
		Time:                t.Time,
		Price:               t.Price,
		Category:            t.Category,
	}
}

type MenuDelete struct {
	ID string
}

type MenuRepository interface {
	UpdateMenuItem(ctx context.Context, dto *MenuModel) error
	CreateMenuItem(ctx context.Context, dto *MenuCreate) error
	DeleteMenuItem(ctx context.Context, dto *MenuDelete) error
	GetMenuItems(ctx context.Context) ([]*MenuModel, error)
}

type MenuService interface {
	UpdateMenuItem(ctx context.Context, dto *MenuModel) error
	CreateMenuItem(ctx context.Context, dto *MenuCreate) error
	DeleteMenuItem(ctx context.Context, dto *MenuDelete) error
	GetMenuItems(ctx context.Context) ([]*MenuModel, error)
}
