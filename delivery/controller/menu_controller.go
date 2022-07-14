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

type MenuController struct {
	router     *gin.Engine
	usecase    usecase.MenuUseCaseInterface
	middleware middleware.AuthTokenMiddlewareInterface
	api.BaseApi
}

func NewMenuController(router *gin.Engine, uc usecase.MenuUseCaseInterface, mw middleware.AuthTokenMiddlewareInterface) *MenuController {
	controller := MenuController{
		router:     router,
		usecase:    uc,
		middleware: mw,
	}

	routeMenu := controller.router.Group("/menu")
	routeMenu.Use(mw.RequireToken())
	routeMenu.GET("/:id", controller.getMenuById)
	routeMenu.PUT("/update", controller.updateMenu)
	routeMenu.DELETE("/:id", controller.deleteMenu)
	routeMenu.POST("/register", controller.createMenu)

	return &controller
}

func (mc *MenuController) getMenuById(ctx *gin.Context) {
	var menuFound entity.Menu
	menuID, err := utils.StringToInt64(ctx.Param("id"))
	if err != nil {
		mc.FailedResponse(ctx, utils.WrongInputNumber("id"))
		return
	} else if menuID == 0 {
		mc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	}
	menuFound.ID = uint(menuID)

	err = mc.usecase.GetMenu(&menuFound)
	if err != nil && err.Error() == "record not found" {
		mc.FailedResponse(ctx, utils.DataNotFoundError())
		return
	} else if err != nil {
		mc.FailedResponse(ctx, err)
		return
	}
	mc.SuccessResponse(ctx, menuFound)
}

func (mc *MenuController) updateMenu(ctx *gin.Context) {
	var (
		menuReq     dto.MenuRequest
		updatedMenu entity.Menu
	)
	err := mc.ParseBodyRequest(ctx, &menuReq)
	if menuReq.MenuID == 0 {
		mc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		mc.FailedResponse(ctx, err)
		return
	}

	updatedMenu, err = mc.usecase.UpdateMenu(&menuReq)
	if err != nil {
		mc.FailedResponse(ctx, err)
		return
	}
	mc.SuccessResponse(ctx, updatedMenu)
}

func (mc *MenuController) deleteMenu(ctx *gin.Context) {
	var menuFound entity.Menu
	menuID, err := utils.StringToInt64(ctx.Param("id"))
	if err != nil {
		mc.FailedResponse(ctx, utils.WrongInputNumber("id"))
		return
	} else if menuID == 0 {
		mc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	}
	menuFound.ID = uint(menuID)

	err = mc.usecase.DeleteMenu(&menuFound)
	if err != nil && err.Error() == "record not found" {
		mc.FailedResponse(ctx, utils.DataNotFoundError())
		return
	} else if err != nil {
		mc.FailedResponse(ctx, err)
		return
	}
	mc.SuccessResponse(ctx, menuFound)
}

func (mc *MenuController) createMenu(ctx *gin.Context) {
	var (
		menuReq     dto.MenuRequest
		createdMenu entity.Menu
	)
	err := mc.ParseBodyRequest(ctx, &menuReq)
	if menuReq.MenuName == "" {
		mc.FailedResponse(ctx, utils.RequiredError("menu name"))
		return
	} else if menuReq.MenuPrice == 0 {
		mc.FailedResponse(ctx, utils.RequiredError("menu price"))
	} else if err != nil {
		mc.FailedResponse(ctx, err)
		return
	}

	createdMenu, err = mc.usecase.CreateMenu(&menuReq)
	if err != nil {
		mc.FailedResponse(ctx, err)
		return
	}
	mc.SuccessResponse(ctx, createdMenu)
}
