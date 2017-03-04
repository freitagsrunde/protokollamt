package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotAuthorized() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/protocols")
		c.Abort()
	}
}

func Authorized() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusForbidden)
	}
}
