package config

import (
	"github.com/freitagsrunde/protokollamt/handlers"
	"github.com/freitagsrunde/protokollamt/middleware"
	"github.com/freitagsrunde/protokollamt/models"
	"github.com/gin-gonic/gin"
)

// DefineRoutes initializes a new default gin
// router, defines the provided HTTP endpoints
// of protokollamt, and declared where static
// and template resources are located.
func DefineRoutes(c *models.Config) *gin.Engine {

	// Depending on configured stage this application
	// is deployed in, set log level of gin.
	if c.DeployStage == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if c.DeployStage == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// First page, login and logout.
	router.GET("/", middleware.NotAuthorized(), handlers.Index())
	router.POST("/", middleware.NotAuthorized(), handlers.IndexLogin(c))
	router.GET("/logout", middleware.Authorized(c), handlers.IndexLogout())

	// Endpoint for listing recent protocols.
	router.GET("/protocols", middleware.Authorized(c), handlers.Protocols(c))

	// Endpoints to add a new protocol.
	router.GET("/protocols/new", middleware.Authorized(c), handlers.ProtocolsNew())
	router.POST("/protocols/new", middleware.Authorized(c), handlers.ProtocolsNewUpload(c))

	// Endpoints to review, change, and
	// publish one specific protocol.
	router.GET("/protocols/view/:id", middleware.Authorized(c), handlers.ProtocolsSingle(c))
	router.POST("/protocols/view/:id", middleware.Authorized(c), handlers.ProtocolsSingleChange())
	router.GET("/protocols/view/:id/reprocess", middleware.Authorized(c), handlers.ProtocolsSingleReprocess())
	router.GET("/protocols/view/:id/publish", middleware.Authorized(c), handlers.ProtocolsSinglePublish())

	// Endpoint for listing existing pipeline steps.
	router.GET("/pipeline", middleware.Authorized(c), handlers.Pipeline(c))

	// Endpoints for manipulating removal steps in
	// analyze pipeline of protocols.
	router.POST("/pipeline/removals/add", middleware.Authorized(c), handlers.PipelineRemovalsAdd(c))
	router.DELETE("/pipeline/removals/view/:id", middleware.Authorized(c), handlers.PipelineRemovalsDelete(c))

	// Endpoints for manipulating replacement steps
	// in analyze pipeline of protocols.
	router.POST("/pipeline/replacements/add", middleware.Authorized(c), handlers.PipelineReplacementsAdd(c))
	router.DELETE("/pipeline/replacements/view/:id", middleware.Authorized(c), handlers.PipelineReplacementsDelete(c))

	// Endpoint for listing mail sending options
	// when publishing a reviewed protocol.
	router.GET("/mail", middleware.Authorized(c), handlers.Mail())

	// Serve static files and HTML templates.
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	return router
}
