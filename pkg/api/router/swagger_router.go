package router

import (
	"fmt"

	"github.com/igab-dev/gorm-gin-swagger/docs"
	"github.com/igab-dev/gorm-gin-swagger/pkg/api/middleware"
	"github.com/igab-dev/gorm-gin-swagger/pkg/configs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func swaggerHandleRequest() *gin.Engine {

	host := configs.GetDB().Host + ":" + configs.GetServerPort()

	docs.SwaggerInfo.Title = "GGS"
	docs.SwaggerInfo.Description = "Gin, Gorm and Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r := gin.Default()

	//CONVERTE LOGS PARA JSON E SALVA
	r.Use(middleware.JsonLoggerMiddleware("swagger"))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func SwaggerStart() {

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := swaggerHandleRequest()

	err = r.Run(fmt.Sprintf(":%s", configs.GetSwaggerPort()))
	if err != nil {
		fmt.Println("Error on up api server", err)
	} else {
		fmt.Printf("API Server listening")
	}
}
