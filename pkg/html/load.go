package html

import "github.com/gin-gonic/gin"



func Load(router *gin.Engine) {
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}

