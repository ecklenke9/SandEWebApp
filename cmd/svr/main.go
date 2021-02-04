package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// creates a new default gin router
	r := gin.Default()

	// route get is designed to respond back with a 200 and a message of "pong"
	// so long as the application is reachable
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("pong!")
		c.JSON(200, gin.H{"message": "pong"})
	})

	// starts the server and exposes routes to the webs
	r.Run()
}
