package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a Todo
// @Schemes		 http
// @Summary      Create a todo
// @Description  Todo POST
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        body  body  TodoCreate  true  "TodoCreate"
// @Success      200 {object} Todo
// @Router       /todo/ [post]
func NewTodo(c *gin.Context) {

	var todo TodoCreate

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error reading json": err.Error()})
		return
	}

	if err := todoCreateValidation(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Validation Error": err.Error()})
	}

	nc, err := new(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Creation error": err})
		return
	}

	c.JSON(http.StatusOK, nc)
}
