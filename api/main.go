package main

import (
	"github.com/gin-gonic/gin"
	"main/api/database"
	"main/api/routes"
)

func main() {
	database.ConnectDatabase()
	r := gin.Default()
	r = routes.SetupRoutes(r)
	err := r.Run()
	if err != nil {
		return
	}
}
