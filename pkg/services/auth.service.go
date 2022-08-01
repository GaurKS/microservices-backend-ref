package services

import (
	"fmt"
	"net/http"

	"github.com/GaurKS/backend-palette/pkg/config"
	"github.com/GaurKS/backend-palette/pkg/dtos"
	"github.com/GaurKS/backend-palette/pkg/models"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// var DB *gorm.DB
// var db Handler = New(DB)

func CheckEmail(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	DB := config.Init()
	db := New(DB)
	var newUser dtos.CreateUser
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		fmt.Println(err.Error())
		c.Abort()
		return
	}

	// Hashing password
	if err := newUser.HashPassword(newUser.Password, &newUser); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		fmt.Println(err.Error())
		c.Abort()
		return
	}

	user := models.User{Email: newUser.Email, Password: newUser.Password, ObjectId: GenerateId(14)}
	if result := db.DB.Create(&user); result.Error != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": result.Error.Error(),
			},
		)
		fmt.Println(result.Error)
		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"data": map[string]string{
				"objectId": user.ObjectId,
				"email":    user.Email,
			},
		},
	)
}

func GenerateToken(c *gin.Context) {
	DB := config.Init()
	db := New(DB)
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		c.Abort()
		return
	}

	// check if email exists and password is correct
	record := db.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": record.Error.Error(),
			},
		)
		c.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.IndentedJSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "invalid credentials",
			},
		)
		c.Abort()
		return
	}

	tokenString, err := GenerateJWT(user.Email, user.Name)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		c.Abort()
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"token": tokenString,
		},
	)
}
