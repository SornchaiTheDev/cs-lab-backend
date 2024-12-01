package main

import (
	"fmt"
	"log"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/sqlx"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.NewConfig()

	db := configs.NewDB(config)

	userRepo := sqlx.NewSqlxUserRepository(db)
	userService := services.NewUserService(userRepo)

	app := fiber.New()

	api := app.Group("/api/v1")

	rest.NewAuthRouter(api, config, userService)

	protectedApi := api.Group("/", middleware.ProtectedRouteMiddleware(config.JWTSecret))

	rest.NewAdminRouter(protectedApi, userService)

	port := fmt.Sprintf(":%v", config.Port)

	err := app.Listen(port)
	if err != nil {
		log.Fatal("Error starting server on Port ", port)
	}

}
