package service

import (
	"context"
	"hair-studio-redmond/service/catalog-service/internal/domain"
)

type service struct {
	repo domain.CatalogRepository
}

func NewService(repo domain.CatalogRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetCatalogItems(ctx context.Context) ([]*domain.CatalogModel, error) {
	return s.repo.GetCatalogItems(ctx)
}
func (s *service) UpdateCatalogItem(ctx context.Context, params *domain.CatalogModel) error {
	return s.repo.UpdateCatalogItem(ctx, params)
}
func (s *service) CreateCatalogItem(ctx context.Context, params *domain.CatalogCreate) error {
	return s.repo.CreateCatalogItem(ctx, params)
}
func (s *service) DeleteCatalogItem(ctx context.Context, params *domain.CatalogDelete) error {
	return s.repo.DeleteCatalogItem(ctx, params)
}
