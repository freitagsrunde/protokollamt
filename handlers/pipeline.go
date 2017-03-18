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

// NewReplacementPayload represents the form data
// needed to create a new replacement element in the
// analysis pipeline.
type NewReplacementPayload struct {
	SearchString  string `form:"replacement-search"`
	ReplaceString string `form:"replacement-replace"`
}

// Pipeline provides an overview list of all
// configured removal and replacement elements
// executed during the analysis pipeline.
func Pipeline(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Removals []models.Removal
		dbConner.GetDBConn().Order("\"created\" asc").Find(&Removals)

		var Replacements []models.Replacement
		dbConner.GetDBConn().Order("\"created\" asc").Find(&Replacements)

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

		// Parse removal form data into above defined payload.
		err := c.BindWith(&Payload, binding.FormPost)
		if err != nil {
			c.Redirect(http.StatusFound, "/pipeline")
			c.Abort()
			return
		}

		// Do not allow empty removal elements.
		if Payload.StartTag == "" || Payload.EndTag == "" {
			c.Redirect(http.StatusFound, "/pipeline")
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

		c.Redirect(http.StatusFound, "/pipeline")
	}
}

func PipelineRemovalsDelete(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}

// PipelineReplacementsAdd expects a NewReplacementPayload
// and adds described replacement element to database.
func PipelineReplacementsAdd(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Payload NewReplacementPayload
		var Replacement models.Replacement

		// Parse replacement form data into above defined payload.
		err := c.BindWith(&Payload, binding.FormPost)
		if err != nil {
			c.Redirect(http.StatusFound, "/pipeline")
			c.Abort()
			return
		}

		// Do not allow empty replacement elements.
		if Payload.SearchString == "" || Payload.ReplaceString == "" {
			c.Redirect(http.StatusFound, "/pipeline")
			c.Abort()
			return
		}

		// Fill new replacement element.
		Replacement.ID = uuid.NewV4().String()
		Replacement.Created = time.Now()
		Replacement.SearchString = Payload.SearchString
		Replacement.ReplaceString = Payload.ReplaceString

		// Save replacement element to database.
		dbConner.GetDBConn().Create(&Replacement)

		c.Redirect(http.StatusFound, "/pipeline")
	}
}

func PipelineReplacementsDelete(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"hello": "lol",
		})
	}
}
