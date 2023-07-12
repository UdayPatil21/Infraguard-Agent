package routes

import (
	"infraguard-agent/api"
	"infraguard-agent/middleware"

	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {

	routeGroup := route.Group("/api").Use(middleware.Auth())
	routeGroup.GET("/checkStatus", api.ServerStatus)
	// routeGroup.POST("/command/execute", sendCommand)
	routeGroup.POST("/script/execute", api.ExecuteScript)
	// windows.InitWindowsRoutes(routeGroup)
	// linux.InitLinuxRoutes(routeGroup)
}
