package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/oniikal3/go/hw/04/todos/services/todo"
)

func main() {
	r := gin.Default()
	todos := r.Group("/api/todos")
	{
		todos.GET("/:id", todo.GetHandler)
		todos.POST("/", todo.PostHandler)
		todos.DELETE("/:id", todo.DeleteHandler)
	}
	r.Run(":1234")
}
