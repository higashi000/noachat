package router

import (
	"github.com/higashi000/noachat/noachatclient"
	"github.com/higashi000/noachat/noachatmsg"
	"github.com/higashi000/noachat/ws"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Msg struct {
	Text string `json:"text"`
}

func NewRouter() *echo.Echo {

	e := echo.New()

	ws.InitMelody()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	noachatmsg.InitWebSocketSettings(e)
	noachatclient.SetFileServe(e)

	e.POST("/send", noachatmsg.Send)

	return e
}
