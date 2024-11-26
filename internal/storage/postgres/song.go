package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"song-library-api/internal/core"
	"song-library-api/internal/storage"
)

var _ storage.SongStorage = (*SongStorage)(nil)

type SongStorage struct {
	pool *pgxpool.Pool
}

// FindByFilter implements storage.SongStorage.
func (s *SongStorage) FindByFilter(ctx context.Context, filters core.SongFilters, pagination core.Pagination) ([]core.Song, error) {
	const op = "storage.postgres.song.FindByFilter"

	filtersMap := make(map[string]any)
	if filters.Name != "" {
		filtersMap["name"] = filters.Name
	}

	if filters.Group != "" {
		filtersMap["group"] = filters.Group
	}

	if filters.ReleaseDate != "" {
		filtersMap["release_date"] = filtersMap
	}

	where := "1=1"
	args := make([]any, 2)
	args[0] = pagination.Offset
	args[1] = pagination.Size

	position := 3

	for k, v := range filtersMap {
		where += fmt.Sprintf(` and "%s"=$%d`, k, position)
		args = append(args, v)
		position++
	}

	query := fmt.Sprintf(`select * from songs where %s offset $1 limit $2`, where)

	fmt.Println(query)

	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return []core.Song{}, fmt.Errorf("%s: %w", op, err)
	}

	songs := make([]core.Song, 0)
	for rows.Next() {
		var song core.Song

		err = rows.Scan(&song.ID, &song.Name, &song.Group, &song.Text, &song.ReleaseDate, &song.Link)
		if err != nil {
			return []core.Song{}, fmt.Errorf("%s: %w", op, err)
		}

		songs = append(songs, song)
	}

	return songs, nil

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
	const op = "storage.postgres.song.Save"

	var exists bool
	query := "select( exists( select 1 from songs where id=$1 ) )"
	if err := s.pool.QueryRow(ctx, query, entity.ID).Scan(&exists); err != nil {
		return entity, fmt.Errorf("%s: %w", op, err)
	}

	if !exists {
		query := `insert into songs (id, name, "group", text, link, release_date) values ($1, $2, $3, $4, $5, $6)`
		_, err := s.pool.Exec(ctx, query, entity.ID, entity.Name, entity.Group, entity.Text, entity.Link, entity.ReleaseDate)
		if err != nil {
			return entity, fmt.Errorf("%s: %w", op, err)
		}

		return entity, nil
	}

	// do update

	return entity, nil
}

// SongLyrics implements storage.SongStorage.
func (s *SongStorage) SongLyrics(ctx context.Context, songID uuid.UUID, verse int) ([]string, error) {
	panic("unimplemented")
}
