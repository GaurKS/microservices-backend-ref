package routes

import (
	"github.com/GaurKS/backend-palette/pkg/services"
	"github.com/gin-gonic/gin"
	// "github.com/GaurKS/backend-palette/pkg/controllers"
)

func TodoRouter(r *gin.RouterGroup, h *services.Handler) {
	r.GET("/getall", h.GetAllTodos)
	r.GET("/get/:id", h.GetTodoById)
	r.POST("/create", h.CreateTodo)
}
