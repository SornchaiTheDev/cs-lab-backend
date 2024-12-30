package main

import (
	"fmt"
	"log"

	"github.com/SornchaiTheDev/cs-lab-backend/configs"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/middlewares"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/rest"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/sqlx"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configs.NewConfig()

	db := configs.NewDB(config)

	userRepo := sqlx.NewSqlxUserRepository(db)
	userService := services.NewUserService(userRepo)

	refreshTokenRepo := sqlx.NewSQLxRefreshTokenRepository(db)
	refreshTokenService := services.NewRefreshTokenService(refreshTokenRepo)

	semesterRepo := sqlx.NewSqlxSemesterRepository(db)
	semesterService := services.NewSemesterService(semesterRepo)

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	api := app.Group("/api/v1")

	rest.NewAuthRouter(api, config, userService, refreshTokenService)

	protectedApi := api.Group("/", middlewares.ProtectedRouteMiddleware(config.JWTSecret))

	rest.NewAdminRouter(protectedApi, userService, semesterService)

	port := fmt.Sprintf(":%v", config.Port)

	err := app.Listen(port)
	if err != nil {
		log.Fatal("Error starting server on Port ", port)
	}

}
