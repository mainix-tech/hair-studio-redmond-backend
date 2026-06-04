package db

import (
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func RunDatabaseMigrations(databaseURL string, fs embed.FS) {
	// Must match the exact structural folder path used in the //go:embed statement!
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Fatalf("failed to create migration io/fs driver: %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, databaseURL)
	if err != nil {
		log.Fatalf("failed to initialize migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run up-migrations: %v", err)
	}
	log.Println("Database schemas are fully up to date!")
}
