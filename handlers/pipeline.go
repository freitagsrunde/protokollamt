package handlers

import (
	"time"

	"net/http"

	"github.com/freitagsrunde/protokollamt/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/satori/go.uuid"
)

// NewRemovalPayload represents the form data
// needed to create a new removal element in the
// analysis pipeline.
type NewRemovalPayload struct {
	StartTag string `form:"removal-start"`
	EndTag   string `form:"removal-end"`
}

// Pipeline provides an overview list of all
// configured removal and replacement elements
// executed during the analysis pipeline.
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

// PipelineRemovalsAdd expects a NewRemovalPayload
// and adds described removal element to database.
func PipelineRemovalsAdd(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Payload NewRemovalPayload
		var Removal models.Removal

		// Get currently existing removal elements.
		var Removals []models.Removal
		dbConner.GetDBConn().Order("\"created\" desc").Find(&Removals)

		// Get currently existing replacement elements.
		var Replacements []models.Replacement
		dbConner.GetDBConn().Order("\"created\" desc").Find(&Replacements)

		// Parse removal form data into above defined payload.
		err := c.BindWith(&Payload, binding.FormPost)
		if err != nil {

			c.HTML(http.StatusBadRequest, "pipeline-list.html", gin.H{
				"PageTitle":    "Protokollamt der Freitagsrunde - Analyse-Pipeline",
				"MainTitle":    "Analyse-Pipeline",
				"FatalError":   "Verarbeitungsfehler. Bitte erneut versuchen.",
				"Removals":     Removals,
				"Replacements": Replacements,
			})
			c.Abort()
			return
		}

		// Construct list of errors, if present.
		errs := make(map[string]string)

		if Payload.StartTag == "" {
			errs["Start-Zeichenkette"] = "Bitte folgendes Feld ausfüllen"
		}

		if Payload.EndTag == "" {
			errs["End-Zeichenkette"] = "Bitte folgendes Feld ausfüllen"
		}

		if len(errs) > 0 {

			c.HTML(http.StatusBadRequest, "pipeline-list.html", gin.H{
				"PageTitle":    "Protokollamt der Freitagsrunde - Analyse-Pipeline",
				"MainTitle":    "Analyse-Pipeline",
				"Errors":       errs,
				"Removals":     Removals,
				"Replacements": Replacements,
			})
			c.Abort()
			return
		}

		// Fill new removal element.
		Removal.ID = uuid.NewV4().String()
		Removal.Created = time.Now()
		Removal.StartTag = Payload.StartTag
		Removal.EndTag = Payload.EndTag

		// Save removal element to database.
		dbConner.GetDBConn().Create(&Removal)

		// Get updated list of removal elements.
		dbConner.GetDBConn().Order("\"created\" desc").Find(&Removals)

		c.HTML(http.StatusOK, "pipeline-list.html", gin.H{
			"PageTitle":    "Protokollamt der Freitagsrunde - Analyse-Pipeline",
			"MainTitle":    "Analyse-Pipeline",
			"Removals":     Removals,
			"Replacements": Replacements,
		})
	}
}

func PipelineRemovalsDelete(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func PipelineReplacementsAdd(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

func PipelineReplacementsDelete(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}
