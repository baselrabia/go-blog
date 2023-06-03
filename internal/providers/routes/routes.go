package routes

import (
	homeRoutes "blog/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine){
	homeRoutes.Routes(r)
}