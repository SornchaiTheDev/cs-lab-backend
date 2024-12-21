package rest

import (
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

	adminRouter.Post("/users/oauth", func(c *fiber.Ctx) error {
		var userRequest requests.CreateUser

		err := c.BodyParser(&userRequest)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		user, err := userService.Create(c.Context(), &userRequest)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		return c.JSON(user)
	})

	adminRouter.Post("/users/credential", func(c *fiber.Ctx) error {
		var userRequest requests.CreateCredentialUser

		err := c.BodyParser(&userRequest)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		user, err := userService.Create(c.Context(), &requests.CreateUser{
			Username:    userRequest.Username,
			DisplayName: userRequest.DisplayName,
			Email:       userRequest.Email,
			Roles:       userRequest.Roles,
		})
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		err = userService.SetPassword(c.Context(), user.ID, userRequest.Password)

		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		return c.JSON(user)
	})

	adminRouter.Get("/users/:userID", func(c *fiber.Ctx) error {
		userID := c.Params("userID")

		user, err := userService.GetByID(c.Context(), userID)
		if err != nil {
			return rerror.ERR_INTERNAL_SERVER_ERROR
		}

		return c.JSON(user)
	})
}
