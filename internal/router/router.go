package router

import (
	"path"
	"path/filepath"

	"github.com/ecklenke9/SandEWebApp/handlers"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Router Has a Gin Engine on it called "Engine"
type Router struct {
	Engine *gin.Engine
}

// New creates a new Gin Engine and returns a Router object with the Engine set to the new gin.Engine
func New() Router {
	// creates a new default gin router
	routerEngine := gin.Default()
	routerEngine.Use(CORSMiddleware())

	routerEngine.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./website/index.html")
		} else {
			c.File("./website/" + path.Join(dir, file))
		}
	})
	return Router{Engine: routerEngine}
}

// BuildRoutes is a 'METHOD' off of the Router struct. This function can only be used by a Router object.
// responsible for building all of our desired routes and applying authentication when implemented
func (g Router) BuildRoutes() {
	// route get is designed to respond back wth a 200 and a message of "pong"
	// so long as the application is reachabl
	g.Engine.GET("/ping", func(c *gin.Context) {
		log.Info().Msg("ping")
		c.JSON(200, gin.H{"message": "pong"})
	})

	g.Engine.GET("/todo", handlers.GetTodoListHandler)
	g.Engine.GET("/todo/:id", handlers.GetTodoByIdHandler)
	g.Engine.POST("/todo", handlers.AddTodoHandler)
	g.Engine.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	g.Engine.PUT("/todo", handlers.CompleteTodoHandler)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
