package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const apiPath = "mini"

// GenerateURLHandler handles requests to generate a mini url
type GenerateURLHandler struct {
	Db      *sql.DB
	BaseURL string
}

// GenerateURL is the structure representing the request to generate a mini url
type GenerateURL struct {
	URL string `json:"url"`
}

// Do generates a mini url given a request like the GenerateURL struct
func (h GenerateURLHandler) Do(ctx echo.Context) error {
	input := new(GenerateURL)

	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Issue with request"+err.Error())
	}

	if input.URL == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Please enter a url value")
	}

	input.URL = strings.TrimSpace(input.URL)
	miniURL, err := h.generateMiniURL(input.URL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Problem generating mini url"+err.Error())
	}

	return ctx.JSON(http.StatusOK, miniURL)
}

func (h GenerateURLHandler) generateMiniURL(originalURL string) (string, error) {
	id := generateBase62()
	sqlStatement := `INSERT INTO url (id, original_url) VALUES ($1, $2)`

	if _, err := h.Db.Exec(sqlStatement, id, originalURL); err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("error saving in db id %s url %s", id, originalURL))
	}

	return fmt.Sprintf("%s/%s/%s", h.BaseURL, apiPath, id), nil
}
