package routes

import (
	"net/http"
	"strconv"

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

func UpdateShelter(c *gin.Context) {
	//Obtener ID del refugio
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error obteniendo de conversion de string a int",
		})
		return
	}

	//Obtener la ID de organizacion
	orgId, exist := c.Get("orgId")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No org id fue encontrada",
		})
		return
	}

	//Obtener el refiguo de la ID
	shelter, err := models.GetSingleShelter(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error obteniendo el refugio",
		})
		return
	}

	//Verificar que el que quiere editar sea la misma orgazacion a laque pertenece el refugio
	if shelter.OrganizationId != orgId {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Solo la organizacion dueña puede editar este refugio",
		})
		return
	}

	//Crear un struct placeholders
	var updatedShelter models.Shelter

	//Llenar al placeholder los datos del body
	err = c.ShouldBindJSON(&updatedShelter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error parseando la info del body",
		})
		return
	}

	//Asignar el ID del Refugio original y el createdAt
	updatedShelter.ID = shelter.ID
	updatedShelter.CreatedAt = shelter.CreatedAt

	//Guardar en la base de datos
	updatedShelter, err = updatedShelter.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error actualizando el refugio",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updatedShelter,
	})
}

func DeleteShelter(c *gin.Context) {
	//Obtener el Id del param
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error obteniendo de conversion de string a int",
		})
		return
	}

	orgId, exist := c.Get("orgId")

	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error el id de la org",
		})
		return
	}
	//Encontrar el refiguo con ese ID
	shelter, err := models.GetSingleShelter(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No existe un refugio con este ID",
		})
	}

	// Ver si la orgazacion que esta haciendo la request es diferente a la que es dueña del shelter
	if shelter.OrganizationId != orgId {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Solo la organizacion dueña puede eliminar este refugio",
		})
		return
	}

	// Eliminar
	err = shelter.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error borrando el refugio",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Refugio Eliminado",
	})
}
