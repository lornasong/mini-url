package main

import (
	"github.com/labstack/echo"
	"github.com/lornasong/mini-url/src/handlers"
)

func main() {

	e := echo.New()
	e.File("/miniurl", "public/index.html")
	e.POST("/mini", handlers.Generate())
	e.GET("/mini/:id", handlers.GoTo())
	e.Start(":8000")
}
