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

// PipelineRemovalsDelete removes the specified
// removal element from database, if existent.
func PipelineRemovalsDelete(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Obtain ID of requested removal element.
		removalID := c.Param("id")

		var Removal models.Removal
		dbConner.GetDBConn().Where("\"id\" = ?", removalID).First(&Removal)

		// If no result could be found for provided
		// removal ID, redirect back to pipeline page.
		if Removal.ID == "" {
			c.Redirect(http.StatusFound, "/pipeline")
			c.Abort()
			return
		}

		// Now we are sure that requested removal element
		// exists, therefore we can safely delete it.
		dbConner.GetDBConn().Delete(&Removal)

		c.Redirect(http.StatusFound, "/pipeline")
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

// PipelineReplacementsDelete removes the specified
// replacement element from database, if existent.
func PipelineReplacementsDelete(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Obtain ID of requested replacement element.
		replacementID := c.Param("id")

		var Replacement models.Replacement
		dbConner.GetDBConn().Where("\"id\" = ?", replacementID).First(&Replacement)

		// If no result could be found for provided
		// replacement ID, redirect back to pipeline page.
		if Replacement.ID == "" {
			c.Redirect(http.StatusFound, "/pipeline")
			c.Abort()
			return
		}

		// Now we are sure that requested replacement element
		// exists, therefore we can safely delete it.
		dbConner.GetDBConn().Delete(&Replacement)

		c.Redirect(http.StatusFound, "/pipeline")
	}
}
