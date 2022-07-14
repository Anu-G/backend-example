package controller

import (
	"wmb-rest-api/delivery/api"
	"wmb-rest-api/delivery/middleware"
	"wmb-rest-api/model/dto"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router     *gin.Engine
	usecase    usecase.CustomerUseCaseInterface
	middleware middleware.AuthTokenMiddlewareInterface
	api.BaseApi
}

func NewCustomerController(router *gin.Engine, uc usecase.CustomerUseCaseInterface, mw middleware.AuthTokenMiddlewareInterface) *CustomerController {
	controller := CustomerController{
		router:     router,
		usecase:    uc,
		middleware: mw,
	}

	routeCustomer := controller.router.Group("/customer")
	routeCustomer.Use(mw.RequireToken())
	routeCustomer.GET("/:id", controller.getCustomerById)
	routeCustomer.PUT("/update", controller.updateCustomer)
	routeCustomer.DELETE("/:id", controller.deleteCustomer)

	return &controller
}

func (cc *CustomerController) getCustomerById(ctx *gin.Context) {
	var customerFound entity.Customer
	customerID, err := utils.StringToInt64(ctx.Param("id"))
	if err != nil {
		cc.FailedResponse(ctx, utils.WrongInputNumber("id"))
		return
	} else if customerID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	}
	customerFound.ID = uint(customerID)

	err = cc.usecase.FindByIdPreload(&customerFound)
	if err != nil && err.Error() == "record not found" {
		cc.FailedResponse(ctx, utils.DataNotFoundError())
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, customerFound)
}

func (cc *CustomerController) updateCustomer(ctx *gin.Context) {
	var (
		customerReq     dto.CustomerRequest
		updatedCustomer entity.Customer
	)
	err := cc.ParseBodyRequest(ctx, &customerReq)
	if customerReq.CustomerID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	updatedCustomer, err = cc.usecase.UpdateCustomer(&customerReq)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, updatedCustomer)
}

func (cc *CustomerController) deleteCustomer(ctx *gin.Context) {
	var customerFound entity.Customer
	customerID, err := utils.StringToInt64(ctx.Param("id"))
	if err != nil {
		cc.FailedResponse(ctx, utils.WrongInputNumber("id"))
		return
	} else if customerID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	}
	customerFound.ID = uint(customerID)

	err = cc.usecase.DeleteCustomer(&customerFound)
	if err != nil && err.Error() == "record not found" {
		cc.FailedResponse(ctx, utils.DataNotFoundError())
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, customerFound)
}
