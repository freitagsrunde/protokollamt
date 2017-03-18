package handlers

import (
	"net/http"

	"github.com/freitagsrunde/protokollamt/models"
	"github.com/gin-gonic/gin"
)

func Pipeline(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Removals []models.Removal
		dbConner.GetDBConn().Order("\"created\" desc").Find(&Removals)

		var Replacements []models.Replacement
		dbConner.GetDBConn().Order("\"created\" desc").Find(&Replacements)

		c.HTML(http.StatusOK, "pipeline-list.html", gin.H{
			"PageTitle":    "Protokollamt der Freitagsrunde - Analyse-Pipeline",
			"MainTitle":    "Analyse-Pipeline",
			"Removals":     Removals,
			"Replacements": Replacements,
		})
	}
}

func PipelineRemovalsAdd() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func PipelineRemovalsAddSubmit() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func PipelineRemovalsDelete() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func PipelineReplacementsAdd() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func PipelineReplacementsAddSubmit() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func PipelineReplacementsDelete() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}
