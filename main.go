package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/GaurKS/backend-palette/pkg/config"
	"github.com/GaurKS/backend-palette/pkg/middlewares"
	"github.com/GaurKS/backend-palette/pkg/routes"
	"github.com/GaurKS/backend-palette/pkg/services"
	"github.com/gin-gonic/gin"
)

func main() {

	// Setting variables for app env

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
	// get the port
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}

	router.Run(port)
}

func getPort() (string, error) {
	// the PORT is supplied by Heroku
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
