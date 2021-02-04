package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	// creates a new default gin router
	routerEngine := gin.Default()
	return routerEngine
}

func BuildRoutes(ginny *gin.Engine) {
	// route get is designed to respond back with a 200 and a message of "pong"
	// so long as the application is reachable
	ginny.GET("/ping", func(c *gin.Context) {
		fmt.Println("pong!")
		c.JSON(200, gin.H{"message": "pong"})
	})

}
