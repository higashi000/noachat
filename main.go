package main

import (
	"github.com/higashi000/noachat/router"
)

func main() {
	noachat := router.NewRouter()

	noachat.E.Logger.Fatal(noachat.E.Start(":5000"))
}
