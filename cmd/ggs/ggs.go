package main

import (
	"github.com/igab-dev/gorm-gin-swagger/pkg/api/router"
)

func main() {

	go router.SwaggerStart()

	router.ApiStart()
}
