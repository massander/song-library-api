package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"song-library-api/internal/core"
	"song-library-api/internal/storage"
)

var _ storage.SongStorage = (*SongStorage)(nil)

type SongStorage struct {
	pool *pgxpool.Pool
}

// Delete implements storage.SongStorage.
func (s *SongStorage) Delete(ctx context.Context, entityID uuid.UUID) error {
	panic("unimplemented")
}

// FindAll implements storage.SongStorage.
func (s *SongStorage) FindAll(ctx context.Context) ([]core.Song, error) {
	panic("unimplemented")
}

// FindById implements storage.SongStorage.
func (s *SongStorage) FindById(ctx context.Context, entityID uuid.UUID) (core.Song, error) {
	panic("unimplemented")
}

// Save implements storage.SongStorage.
func (s *SongStorage) Save(ctx context.Context, entity core.Song) (core.Song, error) {
	panic("unimplemented")
}

// SongLyrics implements storage.SongStorage.
func (s *SongStorage) SongLyrics(ctx context.Context, songID uuid.UUID, verse int) ([]string, error) {
	panic("unimplemented")
}
