package database

import (
	"fmt"
	"log"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBinstance() *gorm.DB {
	dbConfig := config.DbConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DbName,
		dbConfig.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed connecting to database!")
	}
	return db
}
