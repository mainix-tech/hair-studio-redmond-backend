package repository

import "context"

type inmemRepository struct{}

func NewInmemRepository() *inmemRepository { // Return the interface, not the concrete struct
	return &inmemRepository{}
}

func (r *inmemRepository) UpdateProfile(ctx context.Context) error {
	return nil
}
