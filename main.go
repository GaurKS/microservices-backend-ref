package main

import (
	"log"
	"os"

	db "github.com/GaurKS/backend-palette/pkg/config"
	"github.com/GaurKS/backend-palette/pkg/middlewares"
	"github.com/GaurKS/backend-palette/pkg/routes"
	"github.com/GaurKS/backend-palette/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Setting variables for app env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection and handler assignment
	DB := db.Init()
	h := services.New(DB)

	// Assign Gin-Gonic routers with middlewares
	router := gin.Default() //TODO: remove auth from create user
	r := router.Group("/api")
	{
		routes.AuthRouter(r.Group("/auth"), &h)
		r.Use(middlewares.Auth())
		routes.UserRouter(r.Group("/user"), &h)
		routes.TodoRouter(r.Group("/todo"), &h)
	}
	router.Run(os.Getenv("LOCAL_PORT"))
}
