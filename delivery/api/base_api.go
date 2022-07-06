package api

import (
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) ParseBodyRequest(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) SuccessResponse(c *gin.Context, data interface{}) {
	JsonResponseSuccessBuilder(c, data).Send()
}

func (b *BaseApi) FailedResponse(c *gin.Context, err error) {
	JsonResponseFailBuilder(c, err).Send()
}
