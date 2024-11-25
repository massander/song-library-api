package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"

	"song-library-api/internal/storage"
)

type Storage struct {
	pool        *pgxpool.Pool
	databaseURL string

	Songs storage.SongStorage
}

func New(databaseURL string) (*Storage, error) {
	const op = "storage.postgres.New"

	ctx := context.TODO()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pingContext, cancelPing := context.WithTimeout(ctx, time.Second*2)
	defer cancelPing()

	if err := pool.Ping(pingContext); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	songs := &SongStorage{
		pool: pool,
	}

	storage := &Storage{
		pool:        pool,
		databaseURL: databaseURL,
		Songs:       songs,
	}

	return storage, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}

func (s *Storage) Migrate(migrationsFolder string) error {
	const op = "storage.postgres.Migrate"

	databaseURL := strings.Replace(s.databaseURL, "postgres", "pgx5", 1)

	fmt.Println(databaseURL)

	sourceURL := fmt.Sprintf("file:/%s", migrationsFolder)

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No new migrations to apply.")
		} else {
			return fmt.Errorf("%s: failed to run migrations: %w", op, err)
		}
	} else {
		fmt.Println("Migration applied successfully.")
	}

	return nil
}
