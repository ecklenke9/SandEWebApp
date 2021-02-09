package router

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/ecklenke9/SandEWebApp/handlers"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	// creates a new default gin outer
	routerEngine := gin.Default()
	routerEngine.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./website/index.html")
		} else {
			c.File("./website/" + path.Join(dir, file))
		}
	})
	return routerEngine
}
func BuildRoutes(ginny *gin.Engine) {
	// route get is designed to respond back wth a 200 and a message of "pong"
	// so long as the application is reachabl
	ginny.GET("/ping", func(c *gin.Context) {
		fmt.Println("pong!")
		c.JSON(200, gin.H{"message": "pong"})
	})

	ginny.GET("/todo", handlers.GetTodoListHandler)
	ginny.GET("/todo/:id", handlers.GetTodoByIdHandler)
	ginny.POST("/todo", handlers.AddTodoHandler)
	ginny.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	ginny.PUT("/todo", handlers.CompleteTodoHandler)
}
