package routes

import (
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		html.Render(c, "modules/home/html/home", http.StatusOK, gin.H{
			"title": "Home page",
		})
	})

	r.GET("/about", func(c *gin.Context) {
		html.Render(c, "modules/home/html/about", http.StatusOK, gin.H{
			"title": "About page",
		})
	})
}
