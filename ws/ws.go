package ws

import "gopkg.in/olahol/melody.v1"

var WS *melody.Melody

func InitMelody() {
	WS = melody.New()
}
