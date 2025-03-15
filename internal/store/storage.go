package store

import (
	"context"
	"database/sql"
)

var ()

type Storage struct {
	URLS interface {
		ValidateChecksum(ctx context.Context, checksum string) (bool, error)
		CreateShortURL(ctx context.Context, long_url string, checksum string) (int64, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		URLS: &URLStore{db},
	}
}
