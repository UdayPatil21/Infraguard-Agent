package routes

import (
	"infraguard-agent/api/linux"
	"infraguard-agent/api/windows"

	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {

	routeGroup := route.Group("/api")
	windows.InitWindowsRoutes(routeGroup)
	linux.InitLinuxRoutes(routeGroup)
}
