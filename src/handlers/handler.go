package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

// GenerateMiniURL is the structure representing the request to generate a mini url
type GenerateMiniURL struct {
	URL string `json:"url"`
}

// Generate generates a mini url given a request like the GenerateMiniURL struct
func Generate() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := new(GenerateMiniURL)

		if err := ctx.Bind(input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Issue with request"+err.Error())
		}

		if input.URL == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Please enter a url value")
		}

		miniURL, err := generateMiniURL(input.URL)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Problem generating mini url"+err.Error())
		}

		return ctx.JSON(http.StatusOK, miniURL)
	}
}

// GoTo redirects a request for a mini url to the original url
func GoTo() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")

		if len(id) == 0 {
			// TODO: redirect to an error page
		}

		originalURL, err := findOriginalURL(id)
		if err != nil {
			// TODO: redirect this to a pretty error page
			return echo.NewHTTPError(http.StatusNotFound, "Unable to find original url for your tiny URL")
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
