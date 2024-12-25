package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"safehaven.com/m/models"
	"safehaven.com/m/utils"
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

// Sign Up
func CreateOrganization(c *gin.Context) {
	var newOrg models.Organization

	err := c.ShouldBindJSON(&newOrg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error parseando la informacion de body",
		})
		return
	}

	org, err := newOrg.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creando la organizacion",
		})
		return
	}

	jwtToken, err := utils.CreateJWT(org.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creando una token valida",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": org,
		"jwt":  jwtToken,
	})
}

func Login(c *gin.Context) {
	var org models.Organization

	err := c.ShouldBindJSON(&org)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error parseando la informacion de body",
		})
		return
	}

	err = org.ValidateCredential()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	jwtToken, err := utils.CreateJWT(org.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": jwtToken})
}
