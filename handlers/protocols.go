package handlers

import (
	"log"
	"strings"
	"time"

	"net/http"

	"github.com/freitagsrunde/protokollamt/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// DBConner defines the functions needed to
// retrieve and update values in the database.
type DBConner interface {
	GetDBConn() *gorm.DB
}

// NewProtocolPayload represents the form data
// needed to upload a new internal meeting protocol.
type NewProtocolPayload struct {
	Date    string `form:"protocol-date"`
	Content string `form:"protocol-content"`
}

// Protocols lists the existing meeting protocols
// and offers a button to upload a new protocol.
func Protocols(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Protocols []models.Protocol
		dbConner.GetDBConn().Find(&Protocols)

		// Create nicely formatted strings for print
		// in template output.
		for i := range Protocols {
			Protocols[i].MeetingDateString = Protocols[i].MeetingDate.Format("02.01.2006")
		}

		c.HTML(http.StatusOK, "protocols-list.html", gin.H{
			"PageTitle": "Protokollamt der Freitagsrunde",
			"MainTitle": "Protokollamt - Übersicht aller Protokolle",
			"Protocols": Protocols,
		})
	}
}

// ProtocolsNew delivers the page including the
// upload form for a new meeting protocol.
func ProtocolsNew() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "protocols-new.html", gin.H{
			"PageTitle": "Protokollamt der Freitagsrunde",
			"MainTitle": "Protokollamt - Neues Protokoll hochladen",
		})
	}
}

// ProtocolsNewUpload processes submitted data
// for uploading a new meeting protocol.
func ProtocolsNewUpload(dbConner DBConner) gin.HandlerFunc {

	return func(c *gin.Context) {

		var Payload NewProtocolPayload
		var Protocol models.Protocol

		// Parse protocol form data into above defined payload.
		err := c.BindWith(&Payload, binding.FormPost)
		if err != nil {

			c.HTML(http.StatusBadRequest, "protocols-new.html", gin.H{
				"PageTitle":  "Protokollamt der Freitagsrunde",
				"MainTitle":  "Protokollamt - Neues Protokoll hochladen",
				"FatalError": "Verarbeitungsfehler. Bitte erneut versuchen.",
			})
			c.Abort()
			return
		}

		// Remove possibly surrounding whitespace.
		Payload.Date = strings.TrimSpace(Payload.Date)
		Payload.Content = strings.TrimSpace(Payload.Content)

		// Construct list of errors, if present.
		errs := make(map[string]string)

		if Payload.Date == "" {
			errs["Datum"] = "Bitte folgendes Feld ausfüllen"
		}

		if Payload.Content == "" {
			errs["Protokoll"] = "Bitte folgendes Feld ausfüllen"
		}

		if len(errs) > 0 {

			c.HTML(http.StatusBadRequest, "protocols-new.html", gin.H{
				"PageTitle": "Protokollamt der Freitagsrunde",
				"MainTitle": "Protokollamt - Neues Protokoll hochladen",
				"Errors":    errs,
			})
			c.Abort()
			return
		}

		// Attempt to parse time from form.
		meetingDate, err := time.Parse("02.01.2006", Payload.Date)
		if err != nil {
			log.Println(err.Error())

			c.HTML(http.StatusBadRequest, "protocols-new.html", gin.H{
				"PageTitle":  "Protokollamt der Freitagsrunde",
				"MainTitle":  "Protokollamt - Neues Protokoll hochladen",
				"FatalError": "Verarbeitungsfehler. Bitte erneut versuchen.",
			})
			c.Abort()
			return
		}

		// Fill new protocol element.
		Protocol.ID = uuid.NewV4().String()
		Protocol.UploadDate = time.Now()
		Protocol.MeetingDate = meetingDate
		Protocol.Category = models.CategoryFreitagssitzung
		Protocol.InternalVersion = Payload.Content
		Protocol.PublicVersion = Payload.Content
		Protocol.Status = models.StatusInReview

		// Save protocol to database.
		dbConner.GetDBConn().Create(&Protocol)

		c.Redirect(http.StatusFound, "/protocols")
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
