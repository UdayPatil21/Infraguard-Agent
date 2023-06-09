package windows

import (
	"infraguard-agent/helpers/logger"
	"infraguard-agent/middleware"
	model "infraguard-agent/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWindowsRoutes(routeGroup *gin.RouterGroup) {

	r := routeGroup.Group("/windows").Use(middleware.Auth())
	r.POST("/send-command", sendCommand)
}

// Run command on instance
func sendCommand(c *gin.Context) {
	logger.Log.Info("IN:sendCommand")
	input := model.RunCommand{}
	err := c.Bind(&input)
	if err != nil {
		logger.Log.Sugar().Errorf("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	out, err := sendCommandService(input)
	if err != nil {
		logger.Log.Sugar().Errorf("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	logger.Log.Info("OUT:sendCommand")
	c.JSON(http.StatusOK, out)
}
