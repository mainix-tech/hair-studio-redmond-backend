package service

import (
	"context"
	"hair-studio-redmond/service/profile-service/internal/domain"
)

type service struct {
	repo domain.ProfileRepository
}

func NewService(repo domain.ProfileRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) UpdateProfile(ctx context.Context, dto *domain.ProfileModel) error {
	return s.repo.UpdateProfile(ctx, dto)
}
