package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oniikal3/go/hw/04/todos/database"
)

func GetHandler(c *gin.Context) {
	database.Open()
	defer database.Close()
	id, _ := strconv.Atoi(c.Param("id"))
	row, err := database.QueryTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	t := Todo{}
	row.Scan(&t.ID, &t.Title, &t.Status)
	fmt.Println(t)

	c.JSON(http.StatusOK, t)
}
