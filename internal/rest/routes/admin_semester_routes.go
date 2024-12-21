package routes

import (
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/middleware"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/rest/rerror"
	"github.com/gofiber/fiber/v2"
)

func NewAdminSemesterRoutes(router fiber.Router, service services.SemesterService) {
	semesterRouter := router.Group("/semesters")

	semesterRouter.Post("/", middleware.ValidateMiddleware(&requests.Semester{}), func(c *fiber.Ctx) error {
		sem := c.Locals("request").(*requests.Semester)

		createdSem, err := service.Create(c.Context(), sem)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		return c.Status(fiber.StatusCreated).JSON(createdSem)
	})

}
