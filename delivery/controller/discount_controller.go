package controller

import (
	"wmb-rest-api/delivery/api"
	"wmb-rest-api/delivery/middleware"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type DiscountController struct {
	router     *gin.Engine
	usecase    usecase.DiscountUseCaseInterface
	middleware middleware.AuthTokenMiddlewareInterface
	api.BaseApi
}

func NewDiscountController(router *gin.Engine, uc usecase.DiscountUseCaseInterface, mw middleware.AuthTokenMiddlewareInterface) *DiscountController {
	controller := DiscountController{
		router:     router,
		usecase:    uc,
		middleware: mw,
	}

	routeDiscount := controller.router.Group("/discount")
	routeDiscount.Use(mw.RequireToken())
	routeDiscount.GET("/:id", controller.getDiscountById)
	routeDiscount.PUT("/update", controller.updateDiscount)
	routeDiscount.DELETE("/:id", controller.deleteDiscount)
	routeDiscount.POST("/register", controller.createDiscount)

	return &controller
}

func (cc *DiscountController) getDiscountById(ctx *gin.Context) {
	var discountFound entity.Discount
	discountID, err := utils.StringToInt64(ctx.Param("id"))
	if err != nil {
		cc.FailedResponse(ctx, utils.WrongInputNumber("id"))
		return
	} else if discountID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	}
	discountFound.ID = uint(discountID)

	err = cc.usecase.GetDiscountByID(&discountFound)
	if err != nil && err.Error() == "record not found" {
		cc.FailedResponse(ctx, utils.DataNotFoundError())
		return
	} else if err != nil {
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
	var discountFound entity.Discount
	discountID, err := utils.StringToInt64(ctx.Param("id"))
	if err != nil {
		cc.FailedResponse(ctx, utils.WrongInputNumber("id"))
		return
	} else if discountID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	}
	discountFound.ID = uint(discountID)

	err = cc.usecase.DeleteDiscount(&discountFound)
	if err != nil && err.Error() == "record not found" {
		cc.FailedResponse(ctx, utils.DataNotFoundError())
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, discountFound)
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
