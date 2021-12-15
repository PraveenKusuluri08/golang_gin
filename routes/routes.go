package routes

import (
	"github.com/Praveenkusuri08/controller"
	"github.com/gin-gonic/gin"
)

func TodoGetRoute() *gin.Engine {
	r := gin.Default()
	v := r.Group("/api/v1")
	{
		v.GET("todos", controller.GetAllTodos)
		v.POST("todo", controller.CreateTodo)
		v.PUT("todo/:id", controller.UpdateTodo)
	}

	return r
}
