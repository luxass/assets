package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		dt := time.Now()
		return c.String(http.StatusOK, "Hello, World! Rendered from Server: Time: "+dt.String())
	})

	e.GET("/works", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, works!")
	})

	e.ServeHTTP(w, r)
}
