package controller

import (
	"wmb-rest-api/delivery/api"
	"wmb-rest-api/model/dto"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	router  *gin.Engine
	usecase usecase.MenuUseCaseInterface
	api.BaseApi
}

func NewMenuController(router *gin.Engine, uc usecase.MenuUseCaseInterface) *MenuController {
	controller := MenuController{
		router:  router,
		usecase: uc,
	}

	routeMenu := controller.router.Group("/menu")
	routeMenu.GET("/", controller.getMenuById)
	routeMenu.PUT("/", controller.updateMenu)
	routeMenu.DELETE("/", controller.deleteMenu)
	routeMenu.POST("/regsiter", controller.createMenu)

	return &controller
}

func (mc *MenuController) getMenuById(ctx *gin.Context) {
	var menuFound entity.Menu
	err := mc.ParseBodyRequest(ctx, &menuFound)
	if menuFound.ID == 0 {
		mc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		mc.FailedResponse(ctx, err)
		return
	}

	err = mc.usecase.GetMenu(&menuFound)
	if err != nil {
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
	err := mc.ParseBodyRequest(ctx, &menuFound)
	if menuFound.ID == 0 {
		mc.FailedResponse(ctx, utils.RequiredError("id"))
		return
	} else if err != nil {
		mc.FailedResponse(ctx, err)
		return
	}

	err = mc.usecase.DeleteMenu(&menuFound)
	if err != nil {
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
