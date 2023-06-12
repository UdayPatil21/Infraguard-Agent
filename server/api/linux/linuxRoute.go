package linux

import (
	"infraguard-agent/helpers/logger"
	model "infraguard-agent/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitLinuxRoutes(routeGroup *gin.RouterGroup) {
	r := routeGroup.Group("/linux")
	r.POST("/send-command", sendCommand)
	r.POST("/execute-script", executeScript)
	r.POST("/sudo-command", sudoCommand)
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

//Ececute shell script on instance
func executeScript(c *gin.Context) {
	logger.Info("IN:executeScript")
	input := model.Executable{}
	// err := c.Bind(&input)
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	input.Script = data
	out, err := executeScriptService(input)
	if err != nil {
		logger.Error("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	logger.Info("OUT:executeScript")
	c.JSON(http.StatusOK, out)
}

//Run sudo command on instance
func sudoCommand(c *gin.Context) {
	logger.Info("IN:sudoCommand")
	input := model.RunCommand{}
	err := c.Bind(&input)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	out, err := sudoCommandService(input)
	if err != nil {
		logger.Error("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	logger.Info("OUT:sudoCommand")
	c.JSON(http.StatusOK, out)
}
