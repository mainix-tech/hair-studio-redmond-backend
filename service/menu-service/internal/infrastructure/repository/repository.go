package repository

import (
	"context"
	"fmt"
	"hair-studio-redmond/service/menu-service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type postgresRepository struct {
	db *pgx.Conn
}

func NewPostgresRepository(db *pgx.Conn) *postgresRepository {
	return &postgresRepository{db: db}
}

// GetMenuItems 1. GET ALL ITEMS (From previous fix)
func (r *postgresRepository) GetMenuItems(ctx context.Context) ([]*domain.MenuModel, error) {
	query := `
		SELECT id, name, preview_description, detailed_description, time, price, category 
		FROM service;
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query menu items: %w", err)
	}
	defer rows.Close()

	var items []*domain.MenuModel
	for rows.Next() {
		var m domain.MenuModel
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.PreviewDescription,
			&m.DetailedDescription,
			&m.Time,
			&m.Price,
			&m.Category,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan menu item row: %w", err)
		}
		items = append(items, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return items, nil
}

// CreateMenuItem 2. CREATE MENU ITEM
// Accepts domain.MenuCreate (Which explicitly has no ID field)
func (r *postgresRepository) CreateMenuItem(ctx context.Context, params *domain.MenuCreate) error {
	// The database will auto-generate a fresh random UUID for the 'id' column
	query := `
		INSERT INTO service (name, preview_description, detailed_description, time, price, category)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	// Exec is used here because we only care if the query succeeds or fails
	_, err := r.db.Exec(ctx, query,
		params.Name,
		params.PreviewDescription,
		params.DetailedDescription,
		params.Time,
		params.Price,
		params.Category,
	)
	if err != nil {
		return fmt.Errorf("failed to execute insert menu item: %w", err)
	}

	return nil
}

// UpdateMenuItem 3. UPDATE MENU ITEM
// Accepts domain.MenuUpdate (Requires both target ID and replacement fields)
func (r *postgresRepository) UpdateMenuItem(ctx context.Context, params *domain.MenuModel) error {
	query := `
		UPDATE service 
		SET 
			name = $2, 
			preview_description = $3, 
			detailed_description = $4, 
			time = $5, 
			price = $6, 
			category = $7
		WHERE id = $1;
	`

	commandTag, err := r.db.Exec(ctx, query,
		params.ID, // Used to find the matching row target
		params.Name,
		params.PreviewDescription,
		params.DetailedDescription,
		params.Time,
		params.Price,
		params.Category,
	)
	if err != nil {
		return fmt.Errorf("failed to execute update menu item: %w", err)
	}

	// Safety Check: Verify if a row was actually found and updated
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("menu item update failed: no record found with ID %s", params.ID)
	}

	return nil
}

// DeleteMenuItem 4. DELETE MENU ITEM
// Accepts domain.MenuDelete (Strict type mapping for targeting)
func (r *postgresRepository) DeleteMenuItem(ctx context.Context, params *domain.MenuDelete) error {
	query := `DELETE FROM service WHERE id = $1;`

	commandTag, err := r.db.Exec(ctx, query, params.ID)
	if err != nil {
		return fmt.Errorf("failed to execute delete menu item: %w", err)
	}

	// Safety Check: Verify if a row was actually found and destroyed
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("menu item deletion failed: no record found with ID %s", params.ID)
	}

	return nil
}
