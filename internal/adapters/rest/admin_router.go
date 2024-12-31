package rest

import (
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/middlewares"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/rest/routes"
	"github.com/gofiber/fiber/v2"
)

type AdminRouter struct {
	Router          fiber.Router
	UserService     services.UserService
	SemesterService services.SemesterService
	CourseService   services.CourseService
}

func NewAdminRouter(r *AdminRouter) {
	adminRouter := r.Router.Group("/admin", middlewares.AdminMiddleware)

	routes.NewAdminUserRoutes(adminRouter, r.UserService)
	routes.NewAdminSemesterRoutes(adminRouter, r.SemesterService)
	routes.NewAdminCourseRoutes(adminRouter, r.CourseService)
}
