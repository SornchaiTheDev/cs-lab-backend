package main

import (
	"log"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
)

func main() {

	config := configs.NewConfig()

	db := configs.NewDB(config)

	_, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

}
