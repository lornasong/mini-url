package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GenerateMiniURL ... TODO:
type GenerateMiniURL struct {
	URL string `json:"url"`
}

// Generate ... TODO:
func Generate() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := new(GenerateMiniURL)
		if err := ctx.Bind(input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Issue with request"+err.Error())
		}

		miniURL, err := generateMiniURL(input.URL)
		if err != nil {
			// TODO:
		}

		return ctx.JSON(http.StatusOK, miniURL)
	}
}

// GoTo ... TODO:
func GoTo() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		// query for long url based off id
		originalURL, err := findOriginalURL(id)
		if err != nil {
			// TODO:
		}

		return ctx.Redirect(http.StatusSeeOther, originalURL)
	}
}

func findOriginalURL(id string) (string, error) {
	// TODO:
	return "https://www.google.com", nil
}

func generateMiniURL(originalURL string) (string, error) {
	// TODO:
	return "https://www.github.com", nil
}
