package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/seed", seed)
	server.GET("/organizations", GetOrganizations)
	server.GET("/refugees", GetAllRefugees)
	server.GET("/shelter", GetAllShelters)

}
