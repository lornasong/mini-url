package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/lornasong/mini-url/src/handlers"
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

	e := echo.New()
	e.File("/minihome", "public/index.html")
	e.POST("/mini", handlers.GenerateURLHandler{BaseURL: c.BaseURL, APIPath: c.APIPath}.Do)
	e.GET("/mini/:id", handlers.GoToURLHandler{}.Do)
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
