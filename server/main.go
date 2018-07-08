package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	_ "retro-memo/server/docs"
	"retro-memo/server/topic"
)

// @title Retro jar API
// @version 1.0

// @host localhost:8081
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		sprints := v1.Group("/sprints")
		sprints.GET("", func(context *gin.Context) {
			context.String(http.StatusOK, "%s", "HELLO")
		})

		handler := topic.NewHandler()
		topics := v1.Group("/topics")
		topics.GET("/:sprintId", handler.ListTopic)
		topics.POST("", handler.AddTopic)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8081")
}
