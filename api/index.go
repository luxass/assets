package handler

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.GET("/view-source", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "https://github.com/luxass/assets")
	})

	e.GET("/*", func(c echo.Context) error {
		url := c.Request().URL

		if url.Path == "/" {
			url.Path = "/README.md"
		}

		branch := url.Query().Get("branch")
		if branch == "" {
			branch = "main"
		}

		rawURL := "https://raw.githubusercontent.com/luxass/assets/" + branch + url.Path

		resp, err := http.Get(rawURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		c.Response().Header().Set("Content-Type", resp.Header.Get("Content-Type"))

		// set cache control headers
		c.Response().Header().Set("Cache-Control", "public, max-age=3600")

		_, err = io.Copy(c.Response(), resp.Body)
		return err
	})

	e.ServeHTTP(w, r)
}
