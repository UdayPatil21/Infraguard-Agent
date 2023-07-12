package api

import (
	"infraguard-agent/api/linux"
	"infraguard-agent/api/windows"
	"infraguard-agent/helpers/logger"
	model "infraguard-agent/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check server status
func ServerStatus(c *gin.Context) {
	logger.Info("Check Status")
	c.JSON(http.StatusOK, "Success")
}

//Ececute  script on different servers
//On the basis of OS plaform
func ExecuteScript(c *gin.Context) {
	logger.Info("IN:executeScript")

	var input model.Executable
	var response model.CmdOutput
	err := c.Bind(&input)
	if err != nil {
		logger.Error("error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}

	switch input.Platform {
	case "Linux":
		response, err = linux.ExecuteScriptService(input)
		if err != nil {
			logger.Error("Error executing script on linux server", err)
			c.JSON(http.StatusExpectationFailed, err)
		}
	case "Windows":
		response, err = windows.ExecuteScriptService(input)
		if err != nil {
			logger.Error("Error executing script on windows server", err)
			c.JSON(http.StatusExpectationFailed, err)
		}

	case "MAC":

	}

	logger.Info("OUT:executeScript")
	c.JSON(http.StatusOK, response.Output)
}
