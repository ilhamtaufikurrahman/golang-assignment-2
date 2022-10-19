package database

import (
	"assignment-2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	dbPort   = "5432"
	dbName   = "orders_by"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, dbPort, dbName)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Starting DB:", err)
		return
	}

	fmt.Println("Successed connect to DB")
	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
