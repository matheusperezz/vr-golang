package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/api/database"
	"main/api/routes"
)

func main() {
	fmt.Println("Hello, trying to connect to database...")
	database.ConnectDatabase()
	r := gin.Default()
	r = routes.SetupRoutes(r)
	err := r.Run()
	if err != nil {
		return
	}
}
