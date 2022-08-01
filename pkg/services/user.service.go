package services

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/GaurKS/backend-palette/pkg/dtos"
	"github.com/GaurKS/backend-palette/pkg/models"
	"github.com/gin-gonic/gin"
)

func (db Handler) GetAllUser(c *gin.Context) {
	var users []models.User
	if result := db.DB.Find(&users); result.Error != nil {
		fmt.Println(result.Error)
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message":  "SUCCESS",
			"resource": users,
		},
	)
}

// type UserResponse struct {
// 	Name     string
// 	Age      string
// 	Gender   string
// 	Role     string
// 	ImageUrl string
// 	Email    string
// 	// Todos
// }

func (db Handler) GetById(c *gin.Context) {
	var user models.User
	var todos []models.Todo
	id := c.Param("id")
	// rows, err := db.DB.Where("object_id = ?", id).
	// Joins("Join user on user on user.object_id = todo.")
	if err := db.DB.Where("object_id = ?", id).First(&user).Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		c.Abort()
		return
	}
	if err := db.DB.Where("listed_by = ?", id).Find(&todos).Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		c.Abort()
		return
	}
	user.Todos = todos
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"data": user,
		},
	)
}

// TODO: input dto validation
// func (db Handler) CreateUser(c *gin.Context) {
// 	var newUser dtos.CreateUser
// 	if err := c.BindJSON(&newUser); err != nil {
// 		c.IndentedJSON(
// 			http.StatusBadRequest,
// 			gin.H{
// 				"error": err.Error(),
// 			},
// 		)
// 		fmt.Println(err.Error())
// 		c.Abort()
// 		return
// 	}

// 	// Hashing password
// 	if err := newUser.HashPassword(newUser.Password, &newUser); err != nil {
// 		c.IndentedJSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"error": err.Error(),
// 			},
// 		)
// 		fmt.Println(err.Error())
// 		c.Abort()
// 		return
// 	}

// 	user := models.User{Email: newUser.Email, Password: newUser.Password, ObjectId: GenerateId(14)}
// 	if result := db.DB.Create(&user); result.Error != nil {
// 		c.IndentedJSON(
// 			http.StatusInternalServerError,
// 			gin.H{
// 				"error": result.Error,
// 			},
// 		)
// 		fmt.Println(result.Error)
// 		c.Abort()
// 		return
// 	}

// 	c.IndentedJSON(
// 		http.StatusOK,
// 		gin.H{
// 			"data": map[string]string{
// 				"objectId": user.ObjectId,
// 				"email":    user.Email,
// 			},
// 		},
// 	)
// }

// TODO: Protect password and read only field changes
func (db Handler) EditUser(c *gin.Context) {
	var user models.User
	if err := db.DB.Where("object_id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		c.Abort()
		return
	}
	var editUser dtos.EditUser
	if err := c.ShouldBindJSON(&editUser); err != nil {
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

	v := reflect.ValueOf(editUser)
	typeOfV := v.Type()
	inputData := map[string]interface{}{}

	for i := 0; i < v.NumField(); i++ {
		inputData[typeOfV.Field(i).Name] = v.Field(i).Interface()
	}

	if err := db.DB.Model(&user).Updates(inputData).Error; err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(
			http.StatusBadRequest,
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
			"data": user,
		},
	)
}
