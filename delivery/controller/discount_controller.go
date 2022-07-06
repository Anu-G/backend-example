package controller

import (
	"wmb-rest-api/delivery/api"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type DiscountController struct {
	router  *gin.Engine
	usecase usecase.DiscountUseCaseInterface
	api.BaseApi
}

func NewDiscountController(router *gin.Engine, uc usecase.DiscountUseCaseInterface) *DiscountController {
	controller := DiscountController{
		router:  router,
		usecase: uc,
	}

	routeDiscount := controller.router.Group("/discount")
	routeDiscount.GET("/", controller.getDiscountById)
	routeDiscount.PUT("/", controller.updateDiscount)
	routeDiscount.DELETE("/", controller.deleteDiscount)
	routeDiscount.POST("/", controller.createDiscount)

	return &controller
}

func (cc *DiscountController) getDiscountById(ctx *gin.Context) {
	var discountFound entity.Discount
	err := cc.ParseBodyRequest(ctx, &discountFound)
	if discountFound.ID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.GetDiscountByID(&discountFound)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, discountFound)
}

func (cc *DiscountController) updateDiscount(ctx *gin.Context) {
	var discountFound entity.Discount
	err := cc.ParseBodyRequest(ctx, &discountFound)
	if discountFound.ID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.UpdateDiscount(&discountFound)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, discountFound)
}

func (cc *DiscountController) deleteDiscount(ctx *gin.Context) {
	var DiscountFound entity.Discount
	err := cc.ParseBodyRequest(ctx, &DiscountFound)
	if DiscountFound.ID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.DeleteDiscount(&DiscountFound)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, DiscountFound)
}

func (cc *DiscountController) createDiscount(ctx *gin.Context) {
	var newDiscount entity.Discount
	err := cc.ParseBodyRequest(ctx, &newDiscount)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.CreateDiscount(&newDiscount)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, newDiscount)
}
