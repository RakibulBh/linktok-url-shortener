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

func (s *URLStore) CheckChecksum(ctx context.Context, checksum string) (string, error) {
	query := `
		SELECT long_url
		FROM urls
		WHERE checksum = $1
	`

	var redirectUrl string
	err := s.db.QueryRowContext(ctx, query, checksum).Scan(&redirectUrl)
	if err != nil {
		return "", fmt.Errorf("failed to fetch redirect url from db: %w", err)
	}

	return redirectUrl, nil
}

func (s *URLStore) GetRedirectURL(ctx context.Context, rowId int64) (string, error) {
	query := `
		SELECT long_url
		FROM urls
		WHERE id = $1
	`

	// Get the long url from the db
	var longUrl string
	err := s.db.QueryRowContext(ctx, query, rowId).Scan(&longUrl)
	if err != nil {
		return "", err
	}

	return longUrl, nil
}

func (s *URLStore) GetRowID(ctx context.Context, long_url string) (int64, error) {
	query := `
		SELECT id
		FROM urls
		WHERE long_url = $1
	`

	var id int64
	err := s.db.QueryRowContext(ctx, query, long_url).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to get row id: %w", err)
	}

	return id, nil
}
