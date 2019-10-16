package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

// GoToHandler handles requests to go to a mini url
type GoToURLHandler struct{}

// Do redirects a request for a mini url to the original url
func (h GoToURLHandler) Do(ctx echo.Context) error {
	id := ctx.Param("id")

	if len(id) == 0 {
		// TODO: redirect to an error page
	}

	originalURL, err := findOriginalURL(id)
	if err != nil {
		// TODO: redirect this to a pretty error page
		return echo.NewHTTPError(http.StatusNotFound, "Unable to find original url for your mini URL")
	}

	return ctx.Redirect(http.StatusSeeOther, originalURL)
}

func findOriginalURL(id string) (string, error) {
	// TODO: retrieve from database
	return "", errors.New("new error")
}
