package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Protocols(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func ProtocolsNew(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func ProtocolsNewUpload(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func ProtocolsSingle(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func ProtocolsSingleChange(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func ProtocolsSingleReprocess(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func ProtocolsSinglePublish(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}
