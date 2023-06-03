package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/routing"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve() {
	config.Set()
	routing.Init()
	r := routing.GetRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})

	routing.Serve()
}
