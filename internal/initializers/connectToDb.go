package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dbConnection := os.Getenv("DB")

	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	fmt.Println("Successfully connect to db")

	return db
}
