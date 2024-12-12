package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"safehaven.com/m/models"
)

func GetOrganizations(c *gin.Context) {
	orgs, err := models.GetOrgs()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error obteniendo las orgs",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": orgs,
	})
}
