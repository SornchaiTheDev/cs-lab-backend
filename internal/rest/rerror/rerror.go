package rerror

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ERR_BAD_REQUEST           = errors.New("Bad Request")
	ERR_UNAUTHORIZED          = errors.New("Unauthorized")
	ERR_INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
)

func MapErrorWithFiberStatus(err error) int {
	switch err {
	case ERR_BAD_REQUEST:
		return fiber.StatusBadRequest
	case ERR_INTERNAL_SERVER_ERROR:
		return fiber.StatusInternalServerError
	case ERR_UNAUTHORIZED:
		return fiber.StatusUnauthorized
	}

	return fiber.StatusInternalServerError
}
