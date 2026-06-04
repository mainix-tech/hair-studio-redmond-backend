package repository

import (
	"context"
	"hair-studio-redmond/services/profile-service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type postgresRepository struct {
	db *pgx.Conn
}

func NewPostgresRepository(db *pgx.Conn) *postgresRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) UpdateProfile(ctx context.Context, dto *domain.ProfileModel) error {
	query := `UPDATE profiles 
			  SET profile_email = $1,
			      profile_phone = $2,
			      profile_address = $3,
			      profile_title = $4,
			      profile_subtitle = $5
			  WHERE id = $6`

	// Execute your raw native sql using the connection pool
	_, err := r.db.Exec(
		ctx,
		query,
		dto.ProfileEmail,
		dto.ProfilePhone,
		dto.ProfileAddress,
		dto.ProfileTitle,
		dto.ProfileSubtitle,
		dto.ID,
	)
	return err
}
