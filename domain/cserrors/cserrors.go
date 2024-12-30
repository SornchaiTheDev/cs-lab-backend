package cserrors

import (
	"net/http"
	"strconv"
)

type ErrorCode int

func (c *ErrorCode) String() string {
	return strconv.Itoa(int(*c))
}

const (
	BAD_REQUEST           ErrorCode = http.StatusBadRequest
	UNAUTHORIZED          ErrorCode = http.StatusUnauthorized
	INTERNAL_SERVER_ERROR ErrorCode = http.StatusInternalServerError
)

type Error struct {
	Code    ErrorCode
	Message string
}

func New(code ErrorCode, message string) *Error {
	return &Error{Code: code, Message: message}
}

func (c *Error) Error() string {
	return c.Code.String() + "=" + c.Message
}
