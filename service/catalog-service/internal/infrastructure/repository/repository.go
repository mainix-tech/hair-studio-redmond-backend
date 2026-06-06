package repository

import (
	"context"
	"fmt"
	"hair-studio-redmond/service/catalog-service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type postgresRepository struct {
	db *pgx.Conn
}

func NewPostgresRepository(db *pgx.Conn) *postgresRepository {
	return &postgresRepository{db: db}
}

// GetCatalogItems 1. GET ALL ITEMS (From previous fix)
func (r *postgresRepository) GetCatalogItems(ctx context.Context) ([]*domain.CatalogModel, error) {
	query := `
		SELECT id, title, category, img_url
		FROM studio_catalog;
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query catalog items: %w", err)
	}
	defer rows.Close()

	var items []*domain.CatalogModel
	for rows.Next() {
		var m domain.CatalogModel
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Category,
			&m.ImgUrl,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan catalog item row: %w", err)
		}
		items = append(items, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return items, nil
}

// CreateCatalogItem 2. CREATE CATALOG ITEM
// Accepts domain.CatalogCreate (Which explicitly has no ID field)
func (r *postgresRepository) CreateCatalogItem(ctx context.Context, params *domain.CatalogCreate) error {
	// The database will auto-generate a fresh random UUID for the 'id' column
	query := `
		INSERT INTO studio_catalog (title, category, img_url)
		VALUES ($1, $2, $3);
	`

	// Exec is used here because we only care if the query succeeds or fails
	_, err := r.db.Exec(ctx, query,
		params.Title,
		params.Category,
		params.ImgUrl,
	)
	if err != nil {
		return fmt.Errorf("failed to execute insert catalog item: %w", err)
	}

	return nil
}

// UpdateCatalogItem 3. UPDATE CATALOG ITEM
// Accepts domain.CatalogModel (Requires both target ID and replacement fields)
func (r *postgresRepository) UpdateCatalogItem(ctx context.Context, params *domain.CatalogModel) error {
	query := `
		UPDATE studio_catalog 
		SET 
			title = $2, 
			category = $3, 
			img_url = $4
		WHERE id = $1;
	`

	commandTag, err := r.db.Exec(ctx, query,
		params.ID,
		params.Title,
		params.Category,
		params.ImgUrl,
	)
	if err != nil {
		return fmt.Errorf("failed to execute update catalog item: %w", err)
	}

	// Safety Check: Verify if a row was actually found and updated
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("catalog item update failed: no record found with ID %s", params.ID)
	}

	return nil
}

// DeleteCatalogItem 4. DELETE CATALOG ITEM
// Accepts domain.CatalogDelete (Strict type mapping for targeting)
func (r *postgresRepository) DeleteCatalogItem(ctx context.Context, params *domain.CatalogDelete) error {
	query := `DELETE FROM studio_catalog WHERE id = $1;`

	commandTag, err := r.db.Exec(ctx, query, params.ID)
	if err != nil {
		return fmt.Errorf("failed to execute delete catalog item: %w", err)
	}

	// Safety Check: Verify if a row was actually found and destroyed
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("catalog item deletion failed: no record found with ID %s", params.ID)
	}

	return nil
}
