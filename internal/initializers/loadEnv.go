package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := 	godotenv.Load()

	if err != nil {
		log.Fatal("Failed while loading .env file")	
	}
	fmt.Println("Successfully load env variables")
}