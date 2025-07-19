package utils

import (
	"net/http"
)

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Trace   string      `json:"trace"`
	Data    interface{} `json:"data"`
}

func (e *Error) Error() string {
	return e.Message
}

func ErrInternal(message, trace string) error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: message,
		Trace:   trace,
	}
}

func ErrBadRequest(message, trace string) error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: message,
		Trace:   trace,
	}
}

func ErrNotFound(message, trace string) error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: message,
		Trace:   trace,
	}
}

func ErrUnauthorized(message, trace string) error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: message,
		Trace:   trace,
	}
}

func ErrTooManyRequest(message, trace string) error {
	return &Error{
		Code:    http.StatusTooManyRequests,
		Message: message,
		Trace:   trace,
	}
}

func ErrForbidden(message, trace string) error {
	return &Error{
		Code:    http.StatusForbidden,
		Message: message,
		Trace:   trace,
	}
}
