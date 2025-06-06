package router

import (
	"myapp/internal/config"
	"myapp/internal/generated/db"

	"github.com/gin-gonic/gin"

	"myapp/internal/handlers"
)

func SetupRouter(q *db.Queries, cfg *config.Config) *gin.Engine {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)

	engine.LoadHTMLGlob("internal/templates/**/*")

	engine.Static("/static", "internal/st
	api := engine.Group("/api")
	{
		api.POST("/ping", handlers.CreatePingHandler())
	}

	views := engine.Group("/")
	{
		views.GET("/", handlers.RootHandler)
		views.GET("/hello", handlers.HelloHandler)
	}

	return engine
}
