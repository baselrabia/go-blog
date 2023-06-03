package routes

import (
	homeCtrl "blog/internal/modules/home/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	homeController := homeCtrl.New()
	r.GET("/", homeController.Index)
}
