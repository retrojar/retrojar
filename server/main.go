package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	_ "retro-memo/server/docs"
	"retro-memo/server/topic"
)

// @title Retro jar API
// @version 1.0

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := gin.Default()

	conf := &oauth2.Config{
		ClientID:     "1020041870224-8q7qobc8bnhfj3ekuk2uqoqhmsbfqk3g.apps.googleusercontent.com",
		ClientSecret: "iU4RG2RWFnZMqGf_9Hy2B_1d",
		RedirectURL:  "http://localhost:8081/api/v1/login_check",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	v1 := r.Group("/api/v1")
	{
		login := v1.Group("/login")
		login.GET("", func(context *gin.Context) {
			url := conf.AuthCodeURL("MyState1234")
			context.Redirect(http.StatusFound, url)
		})
		v1.GET("/login_check", func(context *gin.Context) {
			token, err := conf.Exchange(oauth2.NoContext, context.Query("code"))
			if err != nil {
				context.AbortWithError(http.StatusBadRequest, err)
				return
			}

			client := conf.Client(oauth2.NoContext, token)
			email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
			if err != nil {
				context.AbortWithError(http.StatusBadRequest, err)
				return
			}
			defer email.Body.Close()
			data, _ := ioutil.ReadAll(email.Body)
			context.String(http.StatusOK, "%s", string(data))
		})

		handler := topic.NewHandler()
		topics := v1.Group("/topics")
		topics.GET("/:sprintId", handler.ListTopic)
		topics.POST("", handler.AddTopic)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", func(context *gin.Context) {
		context.String(200, "%s", "OK")
	})
	r.Run(":8080")
}
