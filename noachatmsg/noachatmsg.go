package noachatmsg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/higashi000/noachat/checkmsg"
	"github.com/higashi000/noachat/ws"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"gopkg.in/olahol/melody.v1"
)

func InitWebSocketSettings(e *echo.Echo) {
	e.GET("/channel/:roomid/ws", func(c echo.Context) error {
		ws.WS.HandleRequest(c.Response().Writer, c.Request())

		return nil
	})

	ws.WS.HandleMessage(func(s *melody.Session, msg []byte) {
		if checkmsg.CheckExclusionWord("./ngword.txt", string(msg)) == nil {
			ws.WS.BroadcastFilter(msg, func(q *melody.Session) bool {
				return q.Request.URL.Path == s.Request.URL.Path
			})
		}
	})

}

func Send(c echo.Context) error {
	var msg Msg
	err := c.Bind(&msg)

	roomID := c.Param("room")
	fmt.Println(roomID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, `{"status": "failed bind msg"}`)
		return errors.Wrap(err, "failed bind msg")
	}

	fmt.Println(checkmsg.CheckExclusionWord("ngword.txt", msg.Text))

	if checkmsg.CheckExclusionWord("./ngword.txt", msg.Text) != nil {
		return c.JSON(http.StatusBadRequest, `{"status":""}`)
	}

	//	ws.WS.Broadcast([]byte(msg.Text))

	ws.WS.BroadcastFilter([]byte(msg.Text), func(q *melody.Session) bool {
		return q.Request.URL.Path == strings.Join([]string{"/channel", roomID, "ws"}, "/")
	})

	c.JSON(http.StatusCreated, msg)
	return nil
}
