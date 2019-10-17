package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// GoToURLHandler handles requests to go to a mini url
type GoToURLHandler struct {
	Db *sql.DB
}

// Do redirects a request for a mini url to the original url
func (h GoToURLHandler) Do(ctx echo.Context) error {
	id := ctx.Param("id")

	if len(id) == 0 {
		// TODO: redirect to an error page
	}

	originalURL, err := h.findOriginalURL(id)
	if err != nil {
		// TODO: redirect this to a pretty error page
		return echo.NewHTTPError(http.StatusNotFound, "Unable to find original url for your mini URL: "+err.Error())
	}

	return ctx.Redirect(http.StatusSeeOther, originalURL)
}

func (h GoToURLHandler) findOriginalURL(id string) (string, error) {
	findURLSQLStatement := `SELECT original_url FROM url WHERE id=$1;`
	row := h.Db.QueryRow(findURLSQLStatement, id)
	var url string
	switch err := row.Scan(&url); err {
	case sql.ErrNoRows:
		return "", errors.Wrap(err, "No url found for "+id)
	case nil:
		return url, nil
	default:
		return "", errors.Wrap(err, "Error looking for id "+id)
	}
}
