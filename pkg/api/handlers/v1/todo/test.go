package todo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateTodoBadRequest(t *testing.T) {

	//mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`

	r := setUpRouter()
	r.POST("/api/v1/todo/", NewTodo)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/api/v1/todo/", bytes.NewBuffer([]byte(`{"title": 123}`)))
	if err != nil {
		t.Errorf("Error in creating request: %v, response: %v", err, req)
	}
	r.ServeHTTP(w, req)

	//responseData, _ := io.ReadAll(w.Body)
	//assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateTodoSuccess(t *testing.T) {

	//mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`

	r := setUpRouter()
	r.POST("/api/v1/todo", NewTodo)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/api/v1/todo", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Errorf("Error in creating request: %v, response: %v", err, req)
	}
	r.ServeHTTP(w, req)

	//responseData, _ := io.ReadAll(w.Body)
	//assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
