package db

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func ConnectPostgres(ctx context.Context, fs embed.FS) (*pgx.Conn, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := "profile_db"

	if user == "" || password == "" || host == "" {
		return nil, fmt.Errorf("database configuration environment variables are missing")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName,
	)

	// 3. Pass the file system bundle directly forward to your migration code!
	log.Println("Checking database schemas and running up-migrations...")
	RunDatabaseMigrations(dsn, fs)

	connectCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(connectCtx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := conn.Ping(connectCtx); err != nil {
		conn.Close(context.Background())
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Printf("Successfully connected to database: %s at %s:%s", dbName, host, port)
	return conn, nil
}
