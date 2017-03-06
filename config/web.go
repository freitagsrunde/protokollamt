package config

import (
	"github.com/freitagsrunde/protokollamt/handlers"
	"github.com/freitagsrunde/protokollamt/middleware"
	"github.com/gin-gonic/gin"
)

// DefineRoutes initializes a new default gin
// router, defines the provided HTTP endpoints
// of protokollamt, and declared where static
// and template resources are located.
func (c *Config) DefineRoutes() *gin.Engine {

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
	router.DELETE("/", middleware.Authorized(c), handlers.IndexLogout())

	// Endpoint for listing recent protocols.
	router.GET("/protocols", middleware.Authorized(c), handlers.Protocols())

	// Endpoints to add a new protocol.
	router.GET("/protocols/new", middleware.Authorized(c), handlers.ProtocolsNew())
	router.POST("/protocols/new", middleware.Authorized(c), handlers.ProtocolsNewUpload())

	// Endpoints to review, change, and
	// publish one specific protocol.
	router.GET("/protocols/view/:id", middleware.Authorized(c), handlers.ProtocolsSingle())
	router.POST("/protocols/view/:id", middleware.Authorized(c), handlers.ProtocolsSingleChange())
	router.GET("/protocols/view/:id/reprocess", middleware.Authorized(c), handlers.ProtocolsSingleReprocess())
	router.GET("/protocols/view/:id/publish", middleware.Authorized(c), handlers.ProtocolsSinglePublish())

	router.GET("/settings", middleware.Authorized(c), handlers.Settings())

	router.GET("/settings/removals/add", middleware.Authorized(c), handlers.SettingsRemovalsAdd())
	router.POST("/settings/removals/add", middleware.Authorized(c), handlers.SettingsRemovalsAddSubmit())
	router.DELETE("/settings/removals/view/:id", middleware.Authorized(c), handlers.SettingsRemovalsDelete())

	router.GET("/settings/replacements/add", middleware.Authorized(c), handlers.SettingsReplacementsAdd())
	router.POST("/settings/replacements/add", middleware.Authorized(c), handlers.SettingsReplacementsAddSubmit())
	router.DELETE("/settings/replacements/view/:id", middleware.Authorized(c), handlers.SettingsReplacementsDelete())

	// Serve static files and HTML templates.
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	return router
}
