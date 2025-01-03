package routes

import (
	"fmt"
	"net/http"
	"strconv"

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

func UpdateRefugee(c *gin.Context) {
	//Obtener id del refugiado
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo parsear el string a int",
		})
		return
	}

	//Buscar refugiado
	refugee, err := models.FindSingleRefugee(id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error encontrando al animalito",
		})
		return
	}
	//TODO: Verificar que el refiguado pertenece a la org?

	//Crear refuiagiado de placeholder
	var auxRefugee models.Refugee

	err = c.ShouldBindJSON(&auxRefugee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se puedo parsear la info del body",
		})
		return
	}

	//Asignar Id y valores que no vienen en el body
	auxRefugee.ID = refugee.ID
	auxRefugee.CreatedAt = refugee.CreatedAt
	auxRefugee.ShelterId = refugee.ShelterId

	//Updatear el refugio
	updatedRefugee, err := auxRefugee.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se puedo actualizar al animalito",
		})
		return
	}

	//Retornar el refugiado actualizado
	c.JSON(http.StatusInternalServerError, gin.H{
		"data": updatedRefugee,
	})
}

func DeleteRefugee(c *gin.Context) {
	//Recuperar ID
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo parsear el string a int",
		})
		return
	}
	//Encontrar el animalito
	refugee, err := models.FindSingleRefugee(id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo encontrar al animalito",
		})
		return
	}

	//Encontrar una manera de verificar que el animalito pertenece a esa orgID

	//Ejecutar la eliminacion
	err = refugee.Delete()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo borrar al animalito",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Animalito eliminado, esperemos que tenga un hogar",
	})
}
