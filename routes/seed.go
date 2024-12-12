package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"safehaven.com/m/models"
)

func seed(c *gin.Context) {
	orgTest := models.Organization{
		Name:        "Test Org",
		Description: "Test description",
	}

	org, err := orgTest.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al guardar la organización",
			"error":   err.Error(),
		})
		return
	}

	shelterTest := models.Shelter{
		Name:           "Patitas seguras",
		Description:    "Adopta no compres",
		Address:        "Cll 23# 157 - 154",
		RefugeeCount:   12,
		ContactPhone:   "+57 123456 7890",
		ContactEmail:   "patitas.seguras@gmail.com",
		OrganizationId: org.ID,
		City:           1,
	}

	shelter, err := shelterTest.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al guardar el refugio",
			"error":   err.Error(),
		})
		return
	}

	refugeeTest := models.Refugee{
		Name:        "Vito",
		RefugeeType: "Perro",
		Size:        "Medium",
		Age:         1,
		AdditionalInfo: json.RawMessage(`{
      "medical_history": "None",
      "languages_spoken": ["English", "Spanish"]
  }`),
		Pictures:  json.RawMessage(`["pic1.jpg", "pic2.jpg"]`),
		ShelterId: shelter.ID,
	}

	refugee, err := refugeeTest.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al guardar el refugiado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Seed completado",
		"organization": org,
		"shelter":      shelter,
		"refugee":      refugee,
	})

}
