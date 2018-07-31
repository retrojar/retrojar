package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "retro-memo/server/docs"
	"retro-memo/server/topic"
	"retro-memo/server/user"
)

// @title Retro jar API
// @version 1.0

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	engine := gin.Default()

	v1 := engine.Group("/api/v1")
	{
		{
			login := v1.Group("/login")
			handler := user.NewHandler()
			login.GET("/google", handler.GetLoginUrl)
			login.GET("/google/jwt", handler.GetJWT)
		}

		{
			topics := v1.Group("/topics")
			handler := topic.NewHandler()
			topics.GET("/:sprintId", handler.ListTopic)
			topics.POST("", handler.AddTopic)
		}
	}

	engine.GET("/health", func(context *gin.Context) {
		context.String(200, "%s", "OK")
	})

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.Run(":8080")
}
