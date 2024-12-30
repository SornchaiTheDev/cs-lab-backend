package routes

import (
	"errors"
	"math"
	"strconv"

	"github.com/SornchaiTheDev/cs-lab-backend/domain/cserrors"
	"github.com/SornchaiTheDev/cs-lab-backend/domain/services"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/adapters/middlewares"
	"github.com/SornchaiTheDev/cs-lab-backend/internal/requests"
	"github.com/gofiber/fiber/v2"
)

func NewAdminSemesterRoutes(router fiber.Router, service services.SemesterService) {
	semesterRouter := router.Group("/semesters")

	semesterRouter.Post("/", middlewares.ValidateMiddleware(&requests.Semester{}), func(c *fiber.Ctx) error {
		sem := c.Locals("request").(*requests.Semester)

		createdSem, err := service.Create(c.Context(), sem)
		if err != nil {
			var csErr *cserrors.Error
			if errors.As(err, &csErr) {
				return err
			}
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error creating semester")
		}

		return c.Status(fiber.StatusCreated).JSON(createdSem)
	})

	semesterRouter.Get("/", func(c *fiber.Ctx) error {
		pageQuery := c.Query("page", "1")
		pageSizeQuery := c.Query("pageSize", "10")
		search := c.Query("search", "")
		sortBy := c.Query("sort_by", "created_at")
		sortOrder := c.Query("sort_order", "desc")

		page, err := strconv.Atoi(pageQuery)
		if err != nil {
			return cserrors.New(cserrors.BAD_REQUEST, "Invalid page")
		}

		pageSize, err := strconv.Atoi(pageSizeQuery)
		if err != nil {
			return cserrors.New(cserrors.BAD_REQUEST, "Invalid page size")
		}

		sems, err := service.GetPagination(c.Context(), page, pageSize, search, sortBy, sortOrder)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error getting semesters")
		}

		count, err := service.Count(c.Context(), search)
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error getting semesters count")
		}

		return c.JSON(fiber.Map{
			"pagination": fiber.Map{
				"page":       page,
				"total_page": math.Ceil(float64(count/pageSize) + 1),
				"total_rows": count,
			},
			"semesters": sems,
		})
	})

	semesterRouter.Get("/:semID", func(c *fiber.Ctx) error {
		semID := c.Params("semID")
		sem, err := service.GetByID(c.Context(), semID)
		if err != nil {
			var csErr *cserrors.Error
			if errors.As(err, &csErr) {
				return err
			}
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error getting semester")
		}

		return c.JSON(sem)
	})

	semesterRouter.Patch("/:semID", func(c *fiber.Ctx) error {
		ID := c.Params("semID")

		var sem requests.Semester

		err := c.BodyParser(&sem)
		if err != nil {
			return cserrors.New(cserrors.BAD_REQUEST, "Error parsing request")
		}

		updateSem, err := service.UpdateByID(c.Context(), ID, &sem)
		if err != nil {
			var csErr *cserrors.Error
			if errors.As(err, &csErr) {
				return err
			}
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error updating semester")
		}

		return c.JSON(updateSem)
	})

	semesterRouter.Delete("/:semID", func(c *fiber.Ctx) error {
		err := service.DeleteByID(c.Context(), c.Params("semID"))
		if err != nil {
			return cserrors.New(cserrors.INTERNAL_SERVER_ERROR, "Error deleting semester")
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

}
