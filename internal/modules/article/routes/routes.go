package routes

import (
	ArticlesCtrl "blog/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller := ArticlesCtrl.New()
	r.GET("/articles/:id", controller.Show)
}
