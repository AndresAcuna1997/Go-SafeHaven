package routes

import (
	"github.com/gin-gonic/gin"
	"safehaven.com/m/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	server.GET("/seed", seed)
	server.GET("/organizations", GetOrganizations)
	server.GET("/refugees", GetAllRefugees)
	server.GET("/shelter", GetAllShelters)

	server.POST("/login", Login)
	server.POST("/signup", CreateOrganization)

	authenticated.POST("/shelter", CreateShelter)
	authenticated.POST("/refugee", CreateRefugee)
	// authenticated.PUT("/organization/:id", UpdateOrg)
	authenticated.PUT("/shelter/:id", UpdateShelter)
	authenticated.PUT("/refugee/:id", UpdateRefugee)

	// authenticated.DELETE("/organization/:id", CreateRefugee)
	authenticated.DELETE("/shelter/:id", DeleteShelter)
	authenticated.DELETE("/refugee/:id", DeleteRefugee)
}
