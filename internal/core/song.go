package core

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Group       string    `json:"group"`
	Text        string    `json:"text"`
	ReleaseDate time.Time `json:"releaseDate"`
	Link        string    `json:"link"`
}
