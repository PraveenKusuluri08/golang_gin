package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Praveenkusuri08/model"
	"github.com/gin-gonic/gin"
)

var todos []model.Todo

func getAllTodos(todos []model.Todo) []model.Todo {
	fmt.Println("Get all todos...")

	return todos
}

func GetAllTodos(c *gin.Context) {
	if data := getAllTodos(todos); data != nil {
		c.JSON(http.StatusOK, data)
	}
}

func createTodo(todo model.Todo) []model.Todo {
	todo.CreatedAt = time.Now().String()
	rand.Seed(time.Now().UnixNano())
	todo.Id = strconv.Itoa(rand.Intn(100))
	todos = append(todos, todo)
	return todos
}

func CreateTodo(c *gin.Context) {
	var todo model.Todo
	c.BindJSON(&todo)
	msg := createTodo(todo)
	c.JSON(200, msg)
}
func updateTodo(id string, todo model.Todo) (string, []model.Todo) {
	for ids, i := range todos {
		fmt.Println("😀", i.Id == id)
		if i.Id != id {
			fmt.Println(todos)
			return "No user exists with the requested id", todos
		} else {
			//removing the data from the slice which matches to the requested id
			todos = append(todos[:ids], todos[ids+1:]...)

			todo.Id = id
			todos = append(todos, todo)
			return "", todos
		}
	}
	return "", todos
}

func UpdateTodo(c *gin.Context) {
	var todo model.Todo
	c.BindJSON(&todo)
	id := c.Params.ByName("id")
	msg, todos := updateTodo(id, todo)

	c.JSON(200, gin.H{"Msg": msg, "todos": todos})
}
