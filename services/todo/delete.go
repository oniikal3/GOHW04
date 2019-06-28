package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oniikal3/go/hw/04/todos/database"
)

func DeleteHandler(c *gin.Context) {
	database.Open()
	defer database.Close()
	id, _ := strconv.Atoi(c.Param("id"))
	err := database.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
	})
}
