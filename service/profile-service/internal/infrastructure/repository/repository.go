package repository

import (
	"context"
	"encoding/json"
	"hair-studio-redmond/service/profile-service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type postgresRepository struct {
	db *pgx.Conn
}

func NewPostgresRepository(db *pgx.Conn) *postgresRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetProfile(ctx context.Context) (*domain.ProfileModel, error) {
	model := &domain.ProfileModel{}

	// 1. Fetch from Home table (use LIMIT 1 since there's only one row)
	err := r.db.QueryRow(ctx, `SELECT title, subtitle, aboutStudioTitle, aboutStudioSubtitle FROM home LIMIT 1`).
		Scan(&model.HomePage.Profile, &model.HomePage.ProfileSubtitle, &model.HomePage.AboutTitle, &model.HomePage.AboutSubtitle)
	if err != nil {
		return nil, err
	}

	// 2. Fetch from Contact table
	var workHoursBytes []byte
	err = r.db.QueryRow(ctx, `SELECT email, phone, address, title, subtitle, work_hours FROM contact LIMIT 1`).
		Scan(&model.ContactPage.ProfileEmail, &model.ContactPage.ProfilePhone, &model.ContactPage.ProfileAddress,
			&model.ContactPage.Title, &model.ContactPage.Subtitle, &workHoursBytes)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(workHoursBytes, &model.ContactPage.WorkHours); err != nil {
		return nil, err
	}

	// 3. Fetch from About table
	err = r.db.QueryRow(ctx, `SELECT title, subtitle, ourStoryTitle, ourStorySubtitle, aboutFounderTitle, aboutFounderSubtitle, aboutFounderReplica FROM about LIMIT 1`).
		Scan(&model.AboutPage.Title, &model.AboutPage.Subtitle, &model.AboutPage.OurStory.Title,
			&model.AboutPage.OurStory.Subtitle, &model.AboutPage.AboutFounder.Title,
			&model.AboutPage.AboutFounder.Subtitle, &model.AboutPage.AboutFounder.Replica)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *postgresRepository) UpdateProfile(ctx context.Context, dto *domain.ProfileModel) error {
	// 1. Start the transaction
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	// Defer a rollback in case of error. If tx.Commit() is called, this does nothing.
	defer tx.Rollback(ctx)

	// 2. Update Contact Table
	// Note: Converting WorkHours to JSON using encoding/json
	workHoursJSON, err := json.Marshal(dto.ContactPage.WorkHours)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `UPDATE contact SET email=$1, phone=$2, address=$3, title=$4, subtitle=$5, work_hours=$6 WHERE id=$7`,
		dto.ContactPage.ProfileEmail, dto.ContactPage.ProfilePhone, dto.ContactPage.ProfileAddress,
		dto.ContactPage.Title, dto.ContactPage.Subtitle, workHoursJSON, dto.ID)
	if err != nil {
		return err
	}

	// 3. Update Home Table
	_, err = tx.Exec(ctx, `UPDATE home SET title=$1, subtitle=$2, aboutStudioTitle=$3, aboutStudioSubtitle=$4 WHERE id=$5`,
		dto.HomePage.Profile, dto.HomePage.ProfileSubtitle, dto.HomePage.AboutTitle, dto.HomePage.AboutSubtitle, dto.ID)
	if err != nil {
		return err
	}

	// 4. Update About Table
	_, err = tx.Exec(ctx, `UPDATE about SET title=$1, subtitle=$2, ourStoryTitle=$3, ourStorySubtitle=$4, 
                                          aboutFounderTitle=$5, aboutFounderSubtitle=$6, aboutFounderReplica=$7 WHERE id=$8`,
		dto.AboutPage.Title, dto.AboutPage.Subtitle, dto.AboutPage.OurStory.Title, dto.AboutPage.OurStory.Subtitle,
		dto.AboutPage.AboutFounder.Title, dto.AboutPage.AboutFounder.Subtitle, dto.AboutPage.AboutFounder.Replica, dto.ID)
	if err != nil {
		return err
	}

	// 5. Commit the transaction
	return tx.Commit(ctx)
}
