package routes

import (
	"github.com/GaurKS/backend-palette/pkg/services"
	"github.com/gin-gonic/gin"
	// "github.com/GaurKS/backend-palette/pkg/controllers"
)

func UserRouter(r *gin.RouterGroup, h *services.Handler) {
	r.GET("/getall", h.GetAllUser)
	r.GET("/get/:id", h.GetById)
	// r.POST("/create", h.CreateUser)
	r.PATCH("/update/:id", h.EditUser)
	// r.POST("/upload/image", services.UploadImage)
}
