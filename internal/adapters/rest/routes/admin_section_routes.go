package routes

import (
	"fmt"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/gofiber/fiber/v2"
)

func NewAdminSectionRoutes(router fiber.Router, service services.SectionService) {
	sectionRouter := router.Group("/sections")

	sectionRouter.Post("/", func(c *fiber.Ctx) error {
		var section requests.Section
		err := c.BodyParser(&section)
		if err != nil {
			return cserrors.New(cserrors.BAD_REQUEST, "Invalid request body")
		}

		createdSection, err := service.Create(c.Context(), &section)
		if err != nil {
			fmt.Println(err)
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error creating section")
		}

		return c.JSON(createdSection)
	})

	sectionRouter.Get("/", func(c *fiber.Ctx) error {
		return nil
	})

	sectionRouter.Get("/:id", func(c *fiber.Ctx) error {
		return nil
	})

	sectionRouter.Put("/:id", func(c *fiber.Ctx) error {
		return nil
	})

	sectionRouter.Delete("/:id", func(c *fiber.Ctx) error {
		return nil
	})
}
