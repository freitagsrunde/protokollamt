package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Settings(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func SettingsRemovalsDelete(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func SettingsReplacementsDelete(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func SettingsRemovalsAdd(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func SettingsRemovalsAddSubmit(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func SettingsReplacementsAdd(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}

func SettingsReplacementsAddSubmit(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "lol",
	})
}
