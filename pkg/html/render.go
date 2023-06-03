package html

import (
	"blog/internal/providers/view"

	"github.com/gin-gonic/gin"
)


func Render(c *gin.Context, name string, code int, data gin.H){
	data = view.WithGlobalData(c, data)
	c.HTML(code, name, data)
}

