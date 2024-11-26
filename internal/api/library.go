package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"song-library-api/internal/core"
)

func (s *APIService) listAllSongs(c *fiber.Ctx) error {
	type Filters struct {
		Name        string `query:"name"`
		Group       string `query:"group"`
		ReleaseDate string `query:"releaseDate"`

		// Есть ли смысл делать филтрацию по тексту и ссылке?
		// Если делать поиск по словам из песни то нужно другую базу данных использовать.
		// Филтрации по ссылки выглядит очень странно
		// Text        string `query:"text"`
		// Link        string `query:"link"`

		Page int `query:"page"`
		Size int `query:"size"`
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
	type Data struct {
		Group string `json:"group"`
		Song  string `json:"song"`
	}

	var data Data
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "INVALID_DATA",
			"message": "",
		})
	}

	// baseURL := os.Getenv("MUSIC_API_URL")

	// URL, err := url.Parse(baseURL)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"code": "INTERNAL_ERROR",
	// 	})
	// }

	// // Add query parameters
	// query := URL.Query()
	// query.Add("name", data.Song)
	// query.Add("group", data.Group)

	// URL.RawQuery = query.Encode()

	// // Create the HTTP request
	// request, err := http.NewRequestWithContext(c.Context(), http.MethodPost, URL.String(), nil)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"code": "INTERNAL_ERROR",
	// 	})
	// }

	// // Add necessary headers
	// request.Header.Set("Content-Type", "application/json")

	// // Create an HTTP client with timeout
	// client := &http.Client{
	// 	Timeout: 10 * time.Second,
	// }

	// response, err := client.Do(request)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"code":    "INTERNAL_ERROR",
	// 		"message": "Failed to make request after retries",
	// 	})
	// }

	// defer response.Body.Close()

	// if response.StatusCode != 200 {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"code": "INTERNAL_ERROR",
	// 	})
	// }

	// type APIResponse struct {
	// 	ReleaseDate time.Time `json:"releaseDate"`
	// 	Text        string    `json:"text"`
	// 	Link        string    `json:"link"`
	// }

	// var apiResponse APIResponse
	// if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"code": "INTERNAL_ERROR",
	// 	})
	// }

	// song := core.Song{
	// 	Name:        data.Song,
	// 	Group:       data.Group,
	// 	Text:        apiResponse.Text,
	// 	Link:        apiResponse.Link,
	// 	ReleaseDate: apiResponse.ReleaseDate,
	// }

	releaseDate, err := time.Parse(time.DateOnly, strings.ReplaceAll("2024.02.02", ".", "-"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": "INTERNAL_ERROR",
		})
	}

	text := "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"

	song := core.Song{
		ID:          uuid.New(),
		Name:        data.Song,
		Group:       data.Group,
		Text:        strings.Split(text, "\n\n"),
		Link:        "https://www.rtretwre.com/watch?v=Xsp3_a-PMTw",
		ReleaseDate: releaseDate,
	}

	song, err = s.storage.Songs.Save(c.Context(), song)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": "INTERNAL_ERROR",
		})
	}

	return c.JSON(song)
}
