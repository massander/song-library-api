package api

import (
	"github.com/gofiber/fiber/v2"
)

type APIService struct {
}

func NewAPIService() *APIService {
	return &APIService{}
}

func (s *APIService) RegisterGateway(router fiber.Router) {
	router.Get("/songs", s.listAllSongs)
	router.Get("/songs/:song_id/lyrics", s.getSongLyrics)
	router.Delete("/songs/:song_id", s.deleteSongFromLibrary)
	router.Patch("songs/:song_id", s.updateSongDetails)
	router.Post("/songs", s.addNewSong)
}
