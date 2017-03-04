package config

import (
	"github.com/freitagsrunde/protokollamt/handlers"
	"github.com/freitagsrunde/protokollamt/middleware"
	"github.com/gin-gonic/gin"
)

func (c *Config) DefineRoutes() *gin.Engine {

	if c.DeployStage == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if c.DeployStage == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/", middleware.NotAuthorized(), handlers.Index)
	router.POST("/", middleware.NotAuthorized(), handlers.IndexLogin)
	router.DELETE("/", middleware.Authorized(), handlers.IndexLogout)

	router.GET("/protocols", middleware.Authorized(), handlers.Protocols)

	router.GET("/protocols/new", middleware.Authorized(), handlers.ProtocolsNew)
	router.POST("/protocols/new", middleware.Authorized(), handlers.ProtocolsNewUpload)

	router.GET("/protocols/view/:id", middleware.Authorized(), handlers.ProtocolsSingle)
	router.POST("/protocols/view/:id", middleware.Authorized(), handlers.ProtocolsSingleChange)
	router.GET("/protocols/view/:id/reprocess", middleware.Authorized(), handlers.ProtocolsSingleReprocess)
	router.GET("/protocols/view/:id/publish", middleware.Authorized(), handlers.ProtocolsSinglePublish)

	router.GET("/settings", middleware.Authorized(), handlers.Settings)

	router.GET("/settings/removals/add", middleware.Authorized(), handlers.SettingsRemovalsAdd)
	router.POST("/settings/removals/add", middleware.Authorized(), handlers.SettingsRemovalsAddSubmit)
	router.DELETE("/settings/removals/view/:id", middleware.Authorized(), handlers.SettingsRemovalsDelete)

	router.GET("/settings/replacements/add", middleware.Authorized(), handlers.SettingsReplacementsAdd)
	router.POST("/settings/replacements/add", middleware.Authorized(), handlers.SettingsReplacementsAddSubmit)
	router.DELETE("/settings/replacements/view/:id", middleware.Authorized(), handlers.SettingsReplacementsDelete)

	return router
}
