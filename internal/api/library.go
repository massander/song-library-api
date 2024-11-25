package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *APIService) listAllSongs(c *fiber.Ctx) error {
	type Filters struct {
		Name        string `query:"name"`
		Group       string `query:"group"`
		Text        string `query:"text"`
		ReleaseDate string `query:"releaseDate"`
		Link        string `query:"link"`
		Page        int    `query:"page"`
	}

	var filters Filters
	if err := c.QueryParser(&filters); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid params",
		})
	}

	return c.JSON(filters)
}

func (s *APIService) getSongLyrics(c *fiber.Ctx) error {
	verse := c.Params("verse")
	return c.JSON(fiber.Map{"verse": verse})
}

func (s *APIService) deleteSongFromLibrary(c *fiber.Ctx) error {
	id := c.Params("song_id")
	songID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "INVALID_PARAMETR",
			"message": "Song ID must be valid UUID",
		})
	}

	return c.JSON(fiber.Map{"id": songID})
}

func (s *APIService) updateSongDetails(c *fiber.Ctx) error {
	id := c.Params("song_id")
	songID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "INVALID_PARAMETR",
			"message": "Song ID must be valid UUID",
		})
	}

	type SongUpdate struct {
	}

	var update SongUpdate

	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "INVALID_DATA",
			"message": "",
		})
	}

	return c.JSON(fiber.Map{"id": songID, "data": update})
}

func (s *APIService) addNewSong(c *fiber.Ctx) error {
	type NewSongRequest struct {
		Group string `json:"group"`
		Song  string `json:"song"`
	}

	var request NewSongRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "INVALID_DATA",
			"message": "",
		})
	}

	return c.JSON(request)
}
