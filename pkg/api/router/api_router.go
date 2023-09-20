package router

import (
	"fmt"

	"github.com/igab-dev/gorm-gin-swagger/pkg/api/handlers/v1/todo"
	"github.com/igab-dev/gorm-gin-swagger/pkg/api/middleware"
	"github.com/igab-dev/gorm-gin-swagger/pkg/configs"

	"github.com/gin-gonic/gin"
)

func apiHandleRequest() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.CORS())
	r.Use(middleware.JsonLoggerMiddleware("app"))

	main := r.Group("api/v1/")
	{
		main.GET("/")
	}

	todos := r.Group("api/v1/todo/")
	{
		todos.POST("/", todo.NewTodo)
		todos.GET("/", todo.AllTodos)
		todos.GET("/:id/", todo.TodoById)
		todos.PUT("/:id/", todo.TodoEdit)
		todos.DELETE("/:id/", todo.DeleteTodo)
	}
	return r
}

func ApiStart() {

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := apiHandleRequest()

	err = r.Run(fmt.Sprintf(":%s", configs.GetServerPort()))
	if err != nil {
		fmt.Println("Error on up api server", err)
	} else {
		fmt.Printf("API Server listening")
	}
}
