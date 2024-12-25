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

func CreateShelter(c *gin.Context) {
	//Crear un shelter
	var newShelter models.Shelter
	//Bind los datos del POST a LA oRG

	err := c.ShouldBindJSON(&newShelter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error obteniendo parseando la informacion del body",
			"err":     err,
		})
		return
	}

	orgId, exists := c.Get("orgId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "orgId no encontrado en el contexto",
		})
		return
	}

	newShelter.OrganizationId = orgId.(int64)

	shelter, err := newShelter.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creando un refugio",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": shelter,
	})
}
