package main

import "github.com/ecklenke9/SandEWebApp/internal/router"

func main() {
	// starts the server and exposes routes to the webs
	r := router.New()
	r.BuildRoutes()
	r.Engine.Run()
}
