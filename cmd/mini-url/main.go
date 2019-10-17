package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/lornasong/mini-url/src/handlers"
	"github.com/pkg/errors"
)

func main() {
	c, err := LoadConfig()
	if err != nil {
		log.Fatalln("Error loading configs", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("$PORT must be set")
	}

	url, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatalln("$DATABASE_URL is required")
	}
	db, err := connect(url)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.File("/minihome", "public/index.html")
	e.POST("/mini", handlers.GenerateURLHandler{Db: db, BaseURL: c.BaseURL, APIPath: c.APIPath}.Do)
	e.GET("/mini/:id", handlers.GoToURLHandler{Db: db}.Do)
	e.Start(":" + port)
}

// Config TODO:
type Config struct {
	BaseURL string `envconfig:"BASE_URL" required:"true`
	APIPath string `envconfig:"API_PATH" required:"true`
}

// LoadConfig loads the configuration object based on the environment and other defaults.
func LoadConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func connect(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "error opening sql db")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "error pinging db")
	}
	return db, nil
}
