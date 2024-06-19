package main

import (
	"github.com/Minz10/REST-API-Go/db"
	"github.com/Minz10/REST-API-Go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost :8080
}

