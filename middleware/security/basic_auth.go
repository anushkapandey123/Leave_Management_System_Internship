package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"main.go/leaves/service"
)

func BasicAuth(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		authParts := strings.SplitN(authHeader, " ", 2)
		if len(authParts) != 2 || authParts[0] != "Basic" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is invalid"})
			c.Abort()
			return
		}

		payload, err := base64.StdEncoding.DecodeString(authParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is invalid"})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is invalid"})
			c.Abort()
			return
		}

		user, err := userService.AuthenticateUser(credentials[0], credentials[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
