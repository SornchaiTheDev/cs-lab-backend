package rest

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

func NewAdminRouter(router fiber.Router, userService services.UserService) {
	adminRouter := router.Group("/admin", middleware.AdminMiddleware)

	adminRouter.Get("/users", func(c *fiber.Ctx) error {
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

		users, err := userService.GetPagination(page, pageSize, search)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		count, err := userService.Count()
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		return c.JSON(fiber.Map{
			"page":       page,
			"total_page": math.Ceil(float64(count/pageSize) + 1),
			"total_rows": count,
			"users":      users,
		})
	})

	adminRouter.Post("/users", func(c *fiber.Ctx) error {
		var userRequest requests.CreateUser

		err := c.BodyParser(&userRequest)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		user, err := userService.Create(c.Context(), &userRequest)
		if err != nil {
			fmt.Println(err)
			return rerror.ERR_INTERNAL_SERVER_ERROR

		}

		return c.JSON(user)
	})
}
