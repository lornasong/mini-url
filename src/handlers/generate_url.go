package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// GenerateURLHandler handles requests to generate a mini url
type GenerateURLHandler struct {
	BaseURL string
	APIPath string
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

	miniURL, err := h.generateMiniURL(input.URL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Problem generating mini url"+err.Error())
	}

	return ctx.JSON(http.StatusOK, miniURL)
}

func (h GenerateURLHandler) generateMiniURL(originalURL string) (string, error) {

	id := generateBase62()
	// TODO: store in persistent database

	return fmt.Sprintf("%s/%s/%s", h.BaseURL, h.APIPath, id), nil
}
