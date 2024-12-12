package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"safehaven.com/m/models"
)

func GetAllShelters(c *gin.Context) {
	shelters, err := models.GetAllShelters()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error obteniendo los refugios",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": shelters,
	})
}
