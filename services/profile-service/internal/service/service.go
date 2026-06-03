package service

import (
	"context"
	"hair-studio-redmond/services/profile-service/internal/domain"
)

type service struct {
	repo domain.ProfileRepository
}

func NewService(repo domain.ProfileRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) UpdateProfile(ctx context.Context) error {
	return s.repo.UpdateProfile(ctx)
}
