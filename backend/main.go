package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Define a simple handler for GET request on '/'
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
