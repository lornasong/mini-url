package main

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/lornasong/mini-url/src/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e := echo.New()
	e.File("/miniurl", "public/index.html")
	e.POST("/mini", handlers.Generate())
	e.GET("/mini/:id", handlers.GoTo())
	e.Start(":" + port)
}
