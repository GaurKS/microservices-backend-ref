package middlewares

import (
	"net/http"

	"github.com/GaurKS/backend-palette/pkg/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.IndentedJSON(
				http.StatusForbidden,
				gin.H{
					"message": "FORBIDDEN REQUEST",
					"error":   "Unauthorized request",
				},
			)
			c.Abort()
			return
		}
		err := services.ValidateToken(tokenString)
		if err != nil {
			c.IndentedJSON(
				http.StatusNotAcceptable,
				gin.H{
					"message": "NOT ACCEPTABLE",
					"error":   err.Error(),
				},
			)
			c.Abort()
			return
		}
		c.Next()
	}
}
