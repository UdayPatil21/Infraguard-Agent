package routes

import (
	"infraguard-agent/api/linux"
	"infraguard-agent/api/windows"
	"infraguard-agent/helpers/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {

	routeGroup := route.Group("/api")
	routeGroup.GET("/checkStatus", serverStatus)
	windows.InitWindowsRoutes(routeGroup)
	linux.InitLinuxRoutes(routeGroup)
}

// Check server status
func serverStatus(c *gin.Context) {
	logger.Info("Check Status")
	c.JSON(http.StatusOK, "Success")
}
