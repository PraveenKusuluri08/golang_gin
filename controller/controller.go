package controller

import (
	"fmt"
	"net/http"

	"github.com/Praveenkusuri08/model"
	"github.com/gin-gonic/gin"
)

func getAllTodos(todos []model.Todo) []model.Todo {
	fmt.Println("Get all todos...")

	return todos
}

func GetAllTodos(c *gin.Context) {
	var todos []model.Todo

	if data := getAllTodos(todos); data != nil {
		c.JSON(http.StatusOK, todos)
	}
}
