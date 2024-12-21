package rest

import (
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/middleware"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/routes"
	"github.com/gofiber/fiber/v2"
)

func NewAdminRouter(router fiber.Router, userService services.UserService) {
	adminRouter := router.Group("/admin", middleware.AdminMiddleware)

	routes.NewAdminUserRoutes(adminRouter, userService)
}
