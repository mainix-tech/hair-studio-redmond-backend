package domain

import "context"

type CatalogModel struct {
	ID       string `bson:"_id"`
	Title    string `bson:"title"`
	Category string `bson:"category"`
	ImgUrl   string `bson:"imgUrl"`
}

type CatalogCreate struct {
	Title    string
	Category string
	ImgUrl   string
}

type CatalogDelete struct {
	ID string
}

type CatalogRepository interface {
	UpdateCatalogItem(ctx context.Context, dto *CatalogModel) error
	CreateCatalogItem(ctx context.Context, dto *CatalogCreate) error
	DeleteCatalogItem(ctx context.Context, dto *CatalogDelete) error
	GetCatalogItems(ctx context.Context) ([]*CatalogModel, error)
}

type CatalogService interface {
	UpdateCatalogItem(ctx context.Context, dto *CatalogModel) error
	CreateCatalogItem(ctx context.Context, dto *CatalogCreate) error
	DeleteCatalogItem(ctx context.Context, dto *CatalogDelete) error
	GetCatalogItems(ctx context.Context) ([]*CatalogModel, error)
}
