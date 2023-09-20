package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Delete a Todo
// @Schemes		 http
// @Summary      Delete a todo
// @Description  Todo DELETE
// @Tags         todo
// @Produce      json
// @Param        id path int true "Todo ID"
// @Success      200 {object} Todo
// @Router       /todo/{id}/ [delete]
func DeleteTodo(c *gin.Context) {

	id := c.Param("id")

	iid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error, id is not a integer": err.Error()})
		return
	}

	if err = delete(iid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error deleting": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Todo "+id+" deleted with success")
}
