package main

import (
	"github.com/gin-gonic/gin"
	"safehaven.com/m/db"
	"safehaven.com/m/routes"
)

func main() {

	server := gin.Default()

	db.Connect()

	routes.RegisterRoutes(server)

	server.Run()

}
