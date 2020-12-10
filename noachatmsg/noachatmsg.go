package noachatmsg

import (
	"net/http"

	"github.com/higashi000/noachat/checkmsg"
	"github.com/higashi000/noachat/ws"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"gopkg.in/olahol/melody.v1"
)

func InitWebSocketSettings(e *echo.Echo) {
	e.GET("/ws", func(c echo.Context) error {
		ws.WS.HandleRequest(c.Response().Writer, c.Request())

		return nil
	})

	ws.WS.HandleMessage(func(s *melody.Session, msg []byte) {
		if checkmsg.CheckExclusionWord("./ngword.txt", string(msg)) == nil {
			ws.WS.Broadcast(msg)
		}
	})
}

func Send(c echo.Context) error {
	var msg Msg
	err := c.Bind(&msg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, `{"status": "failed bind msg"}`)
		return errors.Wrap(err, "failed bind msg")
	}

	if checkmsg.CheckExclusionWord("ngwords.txt", msg.Text) != nil {
		return c.JSON(http.StatusBadRequest, `{"status":""}`)
	}

	ws.WS.Broadcast([]byte(msg.Text))

	c.JSON(http.StatusCreated, msg)
	return nil
}
