package routes

import (
	"fmt"
	"math"
	"strconv"

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

	semesterRouter.Get("/", func(c *fiber.Ctx) error {
		pageQuery := c.Query("page", "1")
		pageSizeQuery := c.Query("pageSize", "10")
		search := c.Query("search", "")

		page, err := strconv.Atoi(pageQuery)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		pageSize, err := strconv.Atoi(pageSizeQuery)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		sems, err := service.GetPagination(c.Context(), page, pageSize, search)
		if err != nil {
			fmt.Println(err)
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		count, err := service.Count(c.Context(), search)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		return c.JSON(fiber.Map{
			"page":       page,
			"total_page": math.Ceil(float64(count/pageSize) + 1),
			"total_rows": count,
			"semesters":  sems,
		})
	})

	semesterRouter.Get("/:semID", func(c *fiber.Ctx) error {
		semID := c.Params("semID")
		sem, err := service.GetByID(c.Context(), semID)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		return c.JSON(sem)
	})

	semesterRouter.Patch("/:semID", func(c *fiber.Ctx) error {
		ID := c.Params("semID")

		var sem requests.Semester

		err := c.BodyParser(&sem)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		updateSem, err := service.UpdateByID(c.Context(), ID, &sem)

		return c.JSON(updateSem)
	})

	semesterRouter.Delete("/:semID", func(c *fiber.Ctx) error {
		err := service.DeleteByID(c.Context(), c.Params("semID"))
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

}
