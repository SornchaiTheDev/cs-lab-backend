package rerror

import (
	"github.com/gofiber/fiber/v2"
)

type RespError string

var (
	ERR_BAD_REQUEST           RespError = "Bad Request"
	ERR_UNAUTHORIZED          RespError = "Unauthorized"
	ERR_INTERNAL_SERVER_ERROR RespError = "Internal Server Error"
)

func MapErrorWithFiberStatus(err RespError) int {
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
