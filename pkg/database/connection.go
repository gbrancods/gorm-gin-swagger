package database

import (
	"fmt"
	"log"

	"github.com/igab-dev/gorm-gin-swagger/pkg/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() (*gorm.DB, error) {
	conf := configs.GetDB()

	cs := fmt.Sprintf("host =%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := gorm.Open(postgres.Open(cs), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}

	return conn, err
}
