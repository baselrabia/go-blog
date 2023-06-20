package routes

import (
	homeRoutes "blog/internal/modules/home/routes"
	articleRoutes "blog/internal/modules/article/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine){
	homeRoutes.Routes(r)
	articleRoutes.Routes(r)
}