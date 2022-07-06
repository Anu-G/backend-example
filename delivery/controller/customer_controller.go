package controller

import (
	"wmb-rest-api/delivery/api"
	"wmb-rest-api/model/dto"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router  *gin.Engine
	usecase usecase.CustomerUseCaseInterface
	api.BaseApi
}

func NewCustomerController(router *gin.Engine, uc usecase.CustomerUseCaseInterface) *CustomerController {
	controller := CustomerController{
		router:  router,
		usecase: uc,
	}

	routeCustomer := controller.router.Group("/customer")
	routeCustomer.GET("/", controller.getCustomerById)
	routeCustomer.PUT("/", controller.updateCustomer)
	routeCustomer.DELETE("/", controller.deleteCustomer)
	routeCustomer.POST("/register", controller.createCustomer)

	return &controller
}

func (cc *CustomerController) getCustomerById(ctx *gin.Context) {
	var customerFound entity.Customer
	err := cc.ParseBodyRequest(ctx, &customerFound)
	if customerFound.ID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.FindByIdPreload(&customerFound)
	if err != nil {
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
	err := cc.ParseBodyRequest(ctx, &customerFound)
	if customerFound.ID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.DeleteCustomer(&customerFound)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, customerFound)
}

func (cc *CustomerController) createCustomer(ctx *gin.Context) {
	var (
		customerReq     dto.CustomerRequest
		createdCustomer entity.Customer
	)
	err := cc.ParseBodyRequest(ctx, &customerReq)
	if customerReq.CustomerName == "" {
		cc.FailedResponse(ctx, utils.RequiredError("customer name"))
		return
	} else if customerReq.MobilePhoneNo == "" {
		cc.FailedResponse(ctx, utils.RequiredError("mobile phone number"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	createdCustomer.MobilePhoneNo = customerReq.MobilePhoneNo
	createdCustomer.CustomerName = customerReq.CustomerName
	createdCustomer, err = cc.usecase.ReadOrCreateCustomer(createdCustomer)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, createdCustomer)
}
