package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Return all Todos
// @Schemes		 http
// @Summary      Return all todos
// @Description  Todo GET
// @Tags         todo
// @Produce      json
// @Success      200 {object} []Todo
// @Router       /todo/ [get]
func AllTodos(c *gin.Context) {

	var todos []Todo
	todos, err := getAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Query error": err})
		return
	}

	c.JSON(http.StatusOK, todos)
}
