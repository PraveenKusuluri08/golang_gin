package routes

import (
	"github.com/Praveenkusuri08/controller"
	"github.com/gin-gonic/gin"
)

func TodoGetRoute() *gin.Engine {
	r := gin.Default()
	v := r.Group("/v1")
	{
		v.GET("todos", controller.GetAllTodos)
	}
	return r
}
