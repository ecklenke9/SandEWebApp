package router

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/YOUR-USER-OR-OG-NAME/YOUR-REPO-NAME/handlers"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	// creates a new default gin outer
	routerEngine := gin.Default()
	routerEngine.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(cRequest.RequestURI)
		ext := filepath.Ext(file) 
		if file == "" || ext == ""{
			c.File("./website/index.html")
		} else {
			c.File("./website/" + path.Join(dir, file))
		}
	})
	return routerEngine


func BuildRoutes(ginny *gin.Engine) {
	// route get is designed to respond back wth a 200 and a message of "pong"
	// so long as the application is reachabl
	ginny.GET("/ping", func(c *gin.Context) {
		fmt.Println("pong!")
		cJSON(200, gin.H{"message": "pong"})
})

	ginny.GET("/todo", handlers.GetTodoListHandlr)
	ginny.POST("/todo", handlers.AddTodoHandler)
	ginny.DELETE("/todo/:id", handlers.DeleteTodoHanler)
	inny.PUT("/todo", handlers.CompleteTodoHandler)
}
