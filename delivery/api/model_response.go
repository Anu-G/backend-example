package api

import (
	"errors"
	"net/http"

	"wmb-rest-api/utils"
)

type Status struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

type Response struct {
	Status
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessMessage(data interface{}) (httpStatusCode int, apiResponse Response) {
	status := Status{
		ResponseCode:    "00",
		ResponseMessage: "success",
	}
	httpStatusCode = http.StatusOK
	apiResponse = Response{
		status, data,
	}
	return
}

func NewErrorMessage(err error) (httpStatusCode int, apiResponse Response) {
	var userError *utils.AppError
	var status Status
	if errors.As(err, &userError) {
		status = Status{
			ResponseCode:    userError.ErrorCode,
			ResponseMessage: userError.ErrorMessage,
		}
		httpStatusCode = userError.ErrorType
	} else {
		status = Status{
			ResponseCode:    "X01",
			ResponseMessage: err.Error(),
		}
		httpStatusCode = http.StatusBadRequest
	}
	apiResponse = Response{
		status, nil,
	}

	return
}
