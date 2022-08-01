package services

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/GaurKS/backend-palette/pkg/dtos"
	"github.com/GaurKS/backend-palette/pkg/models"
	"github.com/gin-gonic/gin"
)

func (db Handler) GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	if result := db.DB.Find(&todos); result.Error != nil {
		fmt.Println(result.Error)
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message":  "SUCCESS",
			"resource": todos,
		},
	)
}

func (db Handler) GetTodoById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := db.DB.Where("object_id = ?", id).First(&todo).Error; err != nil {
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
			"data": todo,
		},
	)
}

func (db Handler) CreateTodo(c *gin.Context) {
	var newTodo dtos.CreateTodo
	if err := c.BindJSON(&newTodo); err != nil {
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

	id := c.GetHeader("UserId")
	fmt.Println("id: ", id)
	newTodo.ListedBy = id
	newTodo.ObjectId = GenerateId(14)

	v := reflect.ValueOf(newTodo)
	typeOfV := v.Type()
	inputData := map[string]interface{}{}

	for i := 0; i < v.NumField(); i++ {
		inputData[typeOfV.Field(i).Name] = v.Field(i).Interface()
	}

	// user := models.Todo{
	// 	ObjectId:       GenerateId(14),
	// 	TodoStatus:     newTodo.TodoStatus,
	// 	ActionText:     newTodo.ActionText,
	// 	Description:    newTodo.Description,
	// 	ListedBy:       newTodo.ListedBy,
	// 	ExpirationDate: newTodo.ExpirationDate,
	// 	ExpirationTime: newTodo.ExpirationTime,
	// 	Assignee:       newTodo.Assignee,
	// }
	var todo models.Todo
	if result := db.DB.Model(&todo).Create(inputData); result.Error != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": result.Error,
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
				"objectId": newTodo.ObjectId,
			},
		},
	)
}
