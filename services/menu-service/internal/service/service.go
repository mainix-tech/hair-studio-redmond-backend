package service

import (
	"context"
	"hair-studio-redmond/services/menu-service/internal/domain"
)

type service struct {
	repo domain.MenuRepository
}

func NewService(repo domain.MenuRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetMenuItems(ctx context.Context) ([]*domain.MenuModel, error) {
	return s.repo.GetMenuItems(ctx)
}
func (s *service) UpdateMenuItem(ctx context.Context, params *domain.MenuModel) error {
	return s.repo.UpdateMenuItem(ctx, params)
}
func (s *service) CreateMenuItem(ctx context.Context, params *domain.MenuCreate) error {
	return s.repo.CreateMenuItem(ctx, params)
}
func (s *service) DeleteMenuItem(ctx context.Context, params *domain.MenuDelete) error {
	return s.repo.DeleteMenuItem(ctx, params)
}
