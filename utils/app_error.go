package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("type: %d, code:%s, err:%s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

func RequiredError(field string) error {
	msg := fmt.Sprintf("%s can't be empty", field)
	return &AppError{
		ErrorCode:    "X02",
		ErrorMessage: msg,
		ErrorType:    http.StatusBadRequest,
	}
}

func DataNotFoundError() error {
	return &AppError{
		ErrorCode:    "X04",
		ErrorMessage: "No Data Found",
		ErrorType:    http.StatusNotFound,
	}
}
