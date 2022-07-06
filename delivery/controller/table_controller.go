package controller

import (
	"wmb-rest-api/delivery/api"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type TableController struct {
	router  *gin.Engine
	usecase usecase.TableUseCaseInterface
	api.BaseApi
}

func NewTableController(router *gin.Engine, uc usecase.TableUseCaseInterface) *TableController {
	controller := TableController{
		router:  router,
		usecase: uc,
	}

	routeTable := controller.router.Group("/table")
	routeTable.GET("/", controller.getTableById)
	routeTable.PUT("/", controller.updateTable)
	routeTable.DELETE("/", controller.deleteTable)
	routeTable.POST("/", controller.createTable)

	return &controller
}

func (cc *TableController) getTableById(ctx *gin.Context) {
	var tableFound entity.Table
	err := cc.ParseBodyRequest(ctx, &tableFound)
	if tableFound.ID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.GetTable(&tableFound)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, tableFound)
}

func (cc *TableController) updateTable(ctx *gin.Context) {
	var tableFound entity.Table
	err := cc.ParseBodyRequest(ctx, &tableFound)
	if tableFound.ID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.UpdateTable(&tableFound, true)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, tableFound)
}

func (cc *TableController) deleteTable(ctx *gin.Context) {
	var tableFound entity.Table
	err := cc.ParseBodyRequest(ctx, &tableFound)
	if tableFound.ID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.DeleteTable(&tableFound)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, tableFound)
}

func (cc *TableController) createTable(ctx *gin.Context) {
	var newTable entity.Table
	err := cc.ParseBodyRequest(ctx, &newTable)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}

	err = cc.usecase.CreateTable(&newTable)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, newTable)
}
