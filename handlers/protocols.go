package handlers

import (
	"net/http"

	"github.com/freitagsrunde/protokollamt/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DBConner defines the functions needed to
// retrieve and update values in the database.
type DBConner interface {
	GetDBConn() *gorm.DB
}

func Protocols(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Protocols []models.Protocol
		dbConner.GetDBConn().Find(&Protocols)

		c.HTML(http.StatusOK, "protocols-list.html", gin.H{
			"PageTitle": "Protokollamt der Freitagsrunde",
			"MainTitle": "Protokollamt",
			"Protocols": Protocols,
		})
	}
}

func ProtocolsNew() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func ProtocolsNewUpload() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func ProtocolsSingle() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func ProtocolsSingleChange() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func ProtocolsSingleReprocess() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func ProtocolsSinglePublish() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}
