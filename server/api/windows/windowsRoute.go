package windows

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWindowsRoutes(routeGroup *gin.RouterGroup) {

	r := routeGroup.Group("/windows")
	r.POST("/send-command", SendCommands)
}

func SendCommands(c *gin.Context) {
	log.Println("Welcome to windows")
	c.JSON(http.StatusOK, "done")
}
