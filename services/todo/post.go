package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oniikal3/go/hw/04/todos/database"
)

func PostHandler(c *gin.Context) {
	database.Open()
	defer database.Close()
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := database.InsertTodo(t.Title, t.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "Success",
		"id":     id,
	})
}
