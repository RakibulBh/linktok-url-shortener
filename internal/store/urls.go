package store

import (
	"context"
	"database/sql"
	"fmt"
)

type URLStore struct {
	db *sql.DB
}

func (s *URLStore) ValidateChecksum(ctx context.Context, checksum string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE checksum = $1)`

	var exists bool
	err := s.db.QueryRowContext(ctx, query, checksum).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to validate checksum: %w", err)
	}
	return exists, nil
}

func (s *URLStore) CreateShortURL(ctx context.Context, long_url string, checksum string) (int64, error) {
	query := `
		INSERT INTO urls (long_url, checksum)
		VALUES ($1, $2)
		RETURNING id
	`

	// Get the ID from the new row
	var id int64
	err := s.db.QueryRowContext(ctx, query, long_url, checksum).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create short url: %w", err)
	}

	return id, nil
}
