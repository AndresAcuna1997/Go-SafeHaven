package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"safehaven.com/m/models"
)

func GetAllRefugees(c *gin.Context) {
	refugees, err := models.GetAllRefugees()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error cargando animalitos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": refugees,
	})
}
