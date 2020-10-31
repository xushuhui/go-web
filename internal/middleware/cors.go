package middleware

import (
	"github.com/gin-gonic/gin"
	"goal/global"
	"net/http"
)

func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", global.ServerSetting.Domain)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE")
		if c.Request.Method == http.MethodOptions {

			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
