package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const stringDeConexao string = "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar banco de dados")
	}
}
