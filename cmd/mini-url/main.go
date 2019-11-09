package main

import (
	"database/sql"
	"log"

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

	db, err := connect(c.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.File("/minihome", "public/index.html")
	e.POST("/mini", handlers.GenerateURLHandler{Db: db, BaseURL: c.BaseURL}.Do)
	e.GET("/mini/:id", handlers.GoToURLHandler{Db: db}.Do)
	e.Start(":" + c.Port)
}

// Config retrieves and store environment variables
type Config struct {
	BaseURL     string `envconfig:"BASE_URL" required:"true`
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true`
	Port        string `required:"true`
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
