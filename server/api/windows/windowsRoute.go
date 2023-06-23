package windows

import (
	"infraguard-agent/helpers/logger"
	model "infraguard-agent/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWindowsRoutes(routeGroup *gin.RouterGroup) {

	r := routeGroup.Group("/platform/windows")
	r.POST("/send-command", sendCommand)
}

//Run command on instance
func sendCommand(c *gin.Context) {
	logger.Info("IN:sendCommand")
	input := model.RunCommand{}
	err := c.Bind(&input)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	out, err := sendCommandService(input)
	if err != nil {
		logger.Error("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	logger.Info("OUT:sendCommand")
	c.JSON(http.StatusOK, out)
}
