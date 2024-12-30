package rest

import (
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/middlewares"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/rest/routes"
	"github.com/gofiber/fiber/v2"
)

func NewAdminRouter(router fiber.Router, userService services.UserService, semesterService services.SemesterService) {
	adminRouter := router.Group("/admin", middlewares.AdminMiddleware)

	routes.NewAdminUserRoutes(adminRouter, userService)
	routes.NewAdminSemesterRoutes(adminRouter, semesterService)
}
