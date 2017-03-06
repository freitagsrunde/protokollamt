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
	router.DELETE("/", middleware.Authorized(), handlers.IndexLogout())

	// Endpoint for listing recent protocols.
	router.GET("/protocols", middleware.Authorized(), handlers.Protocols())

	// Endpoints to add a new protocol.
	router.GET("/protocols/new", middleware.Authorized(), handlers.ProtocolsNew())
	router.POST("/protocols/new", middleware.Authorized(), handlers.ProtocolsNewUpload())

	// Endpoints to review, change, and
	// publish one specific protocol.
	router.GET("/protocols/view/:id", middleware.Authorized(), handlers.ProtocolsSingle())
	router.POST("/protocols/view/:id", middleware.Authorized(), handlers.ProtocolsSingleChange())
	router.GET("/protocols/view/:id/reprocess", middleware.Authorized(), handlers.ProtocolsSingleReprocess())
	router.GET("/protocols/view/:id/publish", middleware.Authorized(), handlers.ProtocolsSinglePublish())

	router.GET("/settings", middleware.Authorized(), handlers.Settings())

	router.GET("/settings/removals/add", middleware.Authorized(), handlers.SettingsRemovalsAdd())
	router.POST("/settings/removals/add", middleware.Authorized(), handlers.SettingsRemovalsAddSubmit())
	router.DELETE("/settings/removals/view/:id", middleware.Authorized(), handlers.SettingsRemovalsDelete())

	router.GET("/settings/replacements/add", middleware.Authorized(), handlers.SettingsReplacementsAdd())
	router.POST("/settings/replacements/add", middleware.Authorized(), handlers.SettingsReplacementsAddSubmit())
	router.DELETE("/settings/replacements/view/:id", middleware.Authorized(), handlers.SettingsReplacementsDelete())

	// Serve static files and HTML templates.
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	return router
}
