package main

import (
	"fmt"
	"log"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.NewConfig()

	_ = configs.NewDB(config)

	app := fiber.New()

	api := app.Group("/api/v1")

	rest.NewAuthRouter(api, config)

	port := fmt.Sprintf(":%v", config.Port)

	err := app.Listen(port)
	if err != nil {
		log.Fatal("Error starting server on Port ", port)
	}

}
