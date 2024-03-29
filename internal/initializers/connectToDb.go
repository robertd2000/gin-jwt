package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
func ConnectToDb() *gorm.DB {
	var err error
	dsn := os.Getenv("DB") // 
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")	
	}
	
	fmt.Println("Successfully connect to db")
	return DB
}