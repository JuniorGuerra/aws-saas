package database

import (
	"app/cmd/emails_sender_service/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

// NewPostgres creates a new PostgreSQL connection using pgx
func NewPostgres(ctx context.Context, conf config.PostgresConfig) (*DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.SSLMode)

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to PostgreSQL database with pgx")
	return &DB{pool: pool}, nil
}

func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) GetPool() *pgxpool.Pool {
	return db.pool
}

func (db *DB) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return db.pool.Query(ctx, sql, args...)
}

func (db *DB) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return db.pool.QueryRow(ctx, sql, args...)
}

func (db *DB) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := db.pool.Exec(ctx, sql, args...)
	return err
}
