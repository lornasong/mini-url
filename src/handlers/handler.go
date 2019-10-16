package handlers

// import (
// 	"errors"
// 	"net/http"

// 	"github.com/labstack/echo"
// )

// // GoTo redirects a request for a mini url to the original url
// func GoTo() echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		id := ctx.Param("id")

// 		if len(id) == 0 {
// 			// TODO: redirect to an error page
// 		}

// 		originalURL, err := findOriginalURL(id)
// 		if err != nil {
// 			// TODO: redirect this to a pretty error page
// 			return echo.NewHTTPError(http.StatusNotFound, "Unable to find original url for your tiny URL")
// 		}

// 		return ctx.Redirect(http.StatusSeeOther, originalURL)
// 	}
// }

// func findOriginalURL(id string) (string, error) {
// 	// TODO:
// 	return "", errors.New("new error")
// 	// return "https://www.google.com", nil
// }
