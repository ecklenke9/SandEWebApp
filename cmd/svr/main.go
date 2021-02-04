package main

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/router"
)

func main() {
	// starts the server and exposes routes to the webs
	r := router.New()
	router.BuildRoutes(r)
	r.Run()
}
