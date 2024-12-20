package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/sqlx"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/middleware"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/rerror"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.NewConfig()

	db := configs.NewDB(config)

	userRepo := sqlx.NewSqlxUserRepository(db)
	userService := services.NewUserService(userRepo)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := rerror.MapErrorWithFiberStatus(err)

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

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
