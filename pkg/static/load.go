package static

import "github.com/gin-gonic/gin"

func Load(r *gin.Engine){
	r.Static("/assets", "./assets")
}