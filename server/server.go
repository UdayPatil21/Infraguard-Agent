package main

import (
	helper "infraguard-agent/helpers"
	"infraguard-agent/helpers/configHelper"
	"infraguard-agent/helpers/logger"
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

	//Add middleware CORS
	r.Use(CORSMiddleware())

	//Init logger
	logger.Init()

	//Init config
	configHelper.InitConfig()

	//Initialize routes
	routes.Init(r)

	err := helper.PreCheck()
	if err != nil {
		StartServer()
	}
	r.Run(":" + configHelper.GetString("Port"))

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept,origin,Cache-Control,X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
