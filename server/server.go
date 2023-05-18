package main

import (
	helper "infraguard-agent/helpers"
	"infraguard-agent/helpers/configHelper"
	"infraguard-agent/helpers/logger"
	"infraguard-agent/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	StartServer()
}

func StartServer() {
	// port := "8080"
	r := gin.Default()

	//Init logger
	logger.Init()
	//Init config
	configHelper.InitConfig()
	//Initialize routes
	routes.Init(r)
	helper.PreCheck()
	r.Run(":" + configHelper.GetString("Port"))

}
