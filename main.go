package main

import (
	"github.com/higashi000/noachat/router"
)

func main() {
	noachat := router.NewRouter()

	noachat.Logger.Fatal(noachat.Start(":5000"))
}
