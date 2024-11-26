package core

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Group       string    `json:"group"`
	Text        []string  `json:"text"`
	ReleaseDate time.Time `json:"releaseDate"`
	Link        string    `json:"link"`
}

type Pagination struct {
	Offset int `query:"offset"`
	Size   int `query:"size"`
}

type SongFilters struct {
	Name        string `query:"name"`
	Group       string `query:"group"`
	ReleaseDate string `query:"releaseDate"`
	// Есть ли смысл делать филтрацию по тексту и ссылке?
	// Если делать поиск по словам из песни то нужно другую базу данных использовать.
	// Филтрации по ссылки выглядит очень странно
	// Text        string `query:"text"`
	// Link        string `query:"link"`
}
