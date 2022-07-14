package controller

import (
	"wmb-rest-api/delivery/api"
	"wmb-rest-api/delivery/middleware"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type TableController struct {
	router     *gin.Engine
	usecase    usecase.TableUseCaseInterface
	middleware middleware.AuthTokenMiddlewareInterface
	api.BaseApi
}

func NewTableController(router *gin.Engine, uc usecase.TableUseCaseInterface, mw middleware.AuthTokenMiddlewareInterface) *TableController {
	controller := TableController{
		router:     router,
		usecase:    uc,
		middleware: mw,
	}

	routeTable := controller.router.Group("/table")
	routeTable.Use(mw.RequireToken())
	routeTable.GET("/:id", controller.getTableById)
	routeTable.PUT("/update", controller.updateTable)
	routeTable.DELETE("/:id", controller.deleteTable)
	routeTable.POST("/register", controller.createTable)

	return &controller
}

func (cc *TableController) getTableById(ctx *gin.Context) {
	var tableFound entity.Table
	tableID, err := utils.StringToInt64(ctx.Param("id"))
	if err != nil {
		cc.FailedResponse(ctx, utils.WrongInputNumber("id"))
		return
	} else if tableID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	}
	tableFound.ID = uint(tableID)

	err = cc.usecase.GetTable(&tableFound)
	if err != nil && err.Error() == "record not found" {
		cc.FailedResponse(ctx, utils.DataNotFoundError())
		return
	} else if err != nil {
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

	err = cc.usecase.UpdateTable(&tableFound)
	if err != nil {
		cc.FailedResponse(ctx, err)
		return
	}
	cc.SuccessResponse(ctx, tableFound)
}

func (cc *TableController) deleteTable(ctx *gin.Context) {
	var tableFound entity.Table
	tableID, err := utils.StringToInt64(ctx.Param("id"))
	if err != nil {
		cc.FailedResponse(ctx, utils.WrongInputNumber("id"))
		return
	} else if tableID == 0 {
		cc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	}
	tableFound.ID = uint(tableID)

	err = cc.usecase.DeleteTable(&tableFound)
	if err != nil && err.Error() == "record not found" {
		cc.FailedResponse(ctx, utils.DataNotFoundError())
		return
	} else if err != nil {
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
