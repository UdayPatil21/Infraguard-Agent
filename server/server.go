package main

import (
	helper "infraguard-agent/helpers"
	"infraguard-agent/helpers/configHelper"
	"infraguard-agent/helpers/logger"
	"infraguard-agent/middleware"
	model "infraguard-agent/models"
	"infraguard-agent/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// Get and set global variables for activation details
	activation_Id := os.Args[1]
	activation_Code := os.Args[2]
	model.Activation_Id = activation_Id
	model.Activation_Code = activation_Code

	StartServer()
}

func StartServer() {
	// port := "8080"
	r := gin.Default()

	//Init config
	configHelper.InitConfig()

	//Add middleware CORS
	r.Use(middleware.CORSMiddleware())

	//Init logger
	logger.Init()

	//Initialize routes
	routes.Init(r)

	err := helper.PreCheck()
	if err != nil {
		StartServer()
	}
	r.Run(":" + configHelper.GetString("Port"))

}
