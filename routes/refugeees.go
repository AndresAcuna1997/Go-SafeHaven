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

func CreateRefugee(c *gin.Context) {
	var newRefugee models.Refugee

	err := c.ShouldBindJSON(&newRefugee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error parseando la informacio ",
		})
		return
	}

	refugee, err := newRefugee.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creando al animalito",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": refugee,
	})

}
