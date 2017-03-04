package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func IndexLogin(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func IndexLogout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}
