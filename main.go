package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"jwt-authentication/database"
	"jwt-authentication/routes"
)

func main() {
	router := gin.Default()
	godotenv.Load()
	database.Main()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*", "https://mywebsite.com"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	routes.AuthRoutes(router)

	router.Run()
}
