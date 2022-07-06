package api

import (
	"github.com/gin-gonic/gin"
)

type AppHttpResponse interface {
	Send()
}

type jsonResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       Response
}

func (j *jsonResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

func JsonResponseSuccessBuilder(c *gin.Context, data interface{}) AppHttpResponse {
	httpStatusCode, resp := NewSuccessMessage(data)
	return &jsonResponse{
		c,
		httpStatusCode,
		resp,
	}
}

func JsonResponseFailBuilder(c *gin.Context, err error) AppHttpResponse {
	httpStatusCode, resp := NewErrorMessage(err)
	return &jsonResponse{
		c,
		httpStatusCode,
		resp,
	}
}

func NewGlobalJsonResponse(c *gin.Context, httpStatusCode int, response Response) AppHttpResponse {
	return &jsonResponse{
		c,
		httpStatusCode,
		response,
	}
}
