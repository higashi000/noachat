package noachatclient

import (
	"net/http"

	"github.com/labstack/echo"
)

func SetFileServe(e *echo.Echo) {

	e.GET("/channel/:roomid", func(c echo.Context) error {
		http.ServeFile(c.Response().Writer, c.Request(), "index.html")

		return nil
	})

	e.GET("/js", func(c echo.Context) error {
		http.ServeFile(c.Response().Writer, c.Request(), "index.js")

		return nil
	})

	e.GET("/favicon", func(c echo.Context) error {
		http.ServeFile(c.Response().Writer, c.Request(), "img/favicon.ico")

		return nil
	})
}
