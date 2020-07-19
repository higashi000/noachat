package router

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/olahol/melody.v1"
)

type NoaChat struct {
	M *melody.Melody
	E *echo.Echo
}

type Msg struct {
	Text string `json:"text"`
}

func NewRouter() NoaChat {
	var msg *Msg

	noachat := NoaChat{
		M: melody.New(),
		E: echo.New(),
	}

	noachat.E.Use(middleware.Logger())
	noachat.E.Use(middleware.Recover())

	noachat.E.GET("/", func(c echo.Context) error {
		http.ServeFile(c.Response().Writer, c.Request(), "index.html")

		return nil
	})

	noachat.E.GET("/ws", func(c echo.Context) error {
		noachat.M.HandleRequest(c.Response().Writer, c.Request())

		return nil
	})

	noachat.E.POST("/send", func(c echo.Context) error {
		err := c.Bind(msg)
		if err != nil {
			return errors.Wrap(err, "failed bind msg")
		}

		noachat.M.Broadcast([]byte(msg.Text))

		return nil
	})

	noachat.M.HandleMessage(func(s *melody.Session, msg []byte) {
		noachat.M.Broadcast(msg)
	})

	return noachat
}
