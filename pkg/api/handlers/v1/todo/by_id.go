package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Return a Todo by ID
// @Schemes		 http
// @Summary      Return a todo by ID
// @Description  Todo GET
// @Tags         todo
// @Produce      json
// @Param        id path int true "Todo ID"
// @Success      200 {object} Todo
// @Router       /todo/{id}/ [get]
func TodoById(c *gin.Context) {

	id := c.Param("id")

	iid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error, todo id is not a integer": err})
		return
	}

	todo, err := byId(iid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Query error": err})
		return
	}

	c.JSON(http.StatusOK, todo)
}
