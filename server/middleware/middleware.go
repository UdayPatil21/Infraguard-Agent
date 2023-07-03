package middleware

import (
	"infraguard-agent/helpers/configHelper"
	"infraguard-agent/middleware/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		//Validate domain
		// hostString := configHelper.GetString("ManagerURL")
		// if !strings.Contains(context.Request.Host, hostString) {
		// 	context.JSON(http.StatusBadRequest, "Unknown Server Connecting .... Error")
		// 	context.Abort()
		// 	return
		// }
		context.Next()
	}
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

func ValidateDomain() gin.HandlerFunc {
	return func(context *gin.Context) {

		hostString := configHelper.GetString("ManagerURL")
		// Check for authorized domain
		if !strings.Contains(context.Request.Host, hostString) {
			context.JSON(http.StatusBadRequest, "Unknown Server Connecting .... Error")
			context.Abort()
			return
		}
		context.Next()
	}
}
