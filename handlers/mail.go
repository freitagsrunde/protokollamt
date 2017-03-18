package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Mail() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}
