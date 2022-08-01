package routes

import (
	"github.com/GaurKS/backend-palette/pkg/services"
	"github.com/gin-gonic/gin"
	// "github.com/GaurKS/backend-palette/pkg/controllers"
)

func AuthRouter(r *gin.RouterGroup, h *services.Handler) {
	r.POST("/validate", services.CheckEmail)
	r.POST("/register", services.CreateUser) //TODO: Handle create user service
	r.POST("/login", services.GenerateToken)
}
