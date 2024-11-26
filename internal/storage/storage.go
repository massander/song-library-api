package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"song-library-api/internal/core"
)

var (
	ErrorSongNotFound = errors.New("song not found")
)

type BaseStorage[T any] interface {
	Save(ctx context.Context, entity T) (T, error)
	FindAll(ctx context.Context) ([]T, error)
	FindById(ctx context.Context, entityID uuid.UUID) (T, error)
	Delete(ctx context.Context, entityID uuid.UUID) error
}

type SongStorage interface {
	BaseStorage[core.Song]
	SongLyrics(ctx context.Context, songID uuid.UUID, verse int) ([]string, error)
	FindByFilter(ctx context.Context, filters core.SongFilters, pagination core.Pagination) ([]core.Song, error)
}
