package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"safehaven.com/m/utils"
)

func Authenticate(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")

	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "A Token is required"})
		return
	}

	orgId, err := utils.VerifyToken(authToken)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Could not verify token"})
		return
	}

	c.Set("orgId", orgId)

	c.Next()
}
