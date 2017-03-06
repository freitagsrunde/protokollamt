package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotAuthorized() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
	}
}

func Authorized() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusForbidden)
	}
}
