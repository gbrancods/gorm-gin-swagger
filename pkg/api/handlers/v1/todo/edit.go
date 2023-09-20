package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Edit a Todo
// @Schemes		 http
// @Summary      Edit a todo
// @Description  Todo PUT
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        id    path  int         true  "Todo ID"
// @Param        body  body  TodoCreate  true  "TodoCreate"
// @Success      200 {object} Todo
// @Router       /todo/{id}/ [post]
func TodoEdit(c *gin.Context) {

	id := c.Param("id")
	var todo Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error reading json": err})
		return
	}

	if err := todoValidation(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Validation Error": err.Error()})
		return
	}

	iid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error, todo id is not a integer": err})
		return
	}

	todo.ID = iid

	ce, err := edit(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error openning database connection": err})
		return
	}

	c.JSON(http.StatusOK, ce)
}
