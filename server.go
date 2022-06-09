package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create instance
	e := echo.New()

	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set route
	e.GET("/", hello)

	// Launch server at port 8003
	e.Logger.Fatal(e.Start(":8003"))
}

// Define handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Golang World!")
}
