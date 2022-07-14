package controller

import (
	"errors"

	"wmb-rest-api/auth"
	"wmb-rest-api/delivery/api"
	"wmb-rest-api/model/dto"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router     *gin.Engine
	customerUC usecase.CustomerUseCaseInterface
	authUC     usecase.AuthUseCaseInterface
	authToken  auth.TokenInterface
	api.BaseApi
}

func NewAuthController(router *gin.Engine, cuc usecase.CustomerUseCaseInterface,
	au usecase.AuthUseCaseInterface, at auth.TokenInterface,
) *AuthController {
	controller := AuthController{
		router:     router,
		customerUC: cuc,
		authUC:     au,
		authToken:  at,
	}

	routerAuth := controller.router.Group("/auth")
	routerAuth.POST("/register", controller.createUserAccount)
	routerAuth.POST("/login", controller.loginUser)

	return &controller
}

func (ac *AuthController) createUserAccount(ctx *gin.Context) {
	var (
		customerReq     dto.RegisterCustomerRequest
		createdCustomer entity.Customer
		createdUser     entity.UserCredential
	)
	err := ac.ParseBodyRequest(ctx, &customerReq)
	if customerReq.UserName == "" {
		ac.FailedResponse(ctx, utils.RequiredError("username"))
		return
	} else if customerReq.UserPassword == "" {
		ac.FailedResponse(ctx, utils.RequiredError("password"))
		return
	} else if customerReq.CustomerName == "" {
		ac.FailedResponse(ctx, utils.RequiredError("customer name"))
		return
	} else if customerReq.MobilePhoneNo == "" {
		ac.FailedResponse(ctx, utils.RequiredError("phone number"))
		return
	} else if customerReq.Email == "" {
		ac.FailedResponse(ctx, utils.RequiredError("email"))
		return
	} else if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	createdCustomer.MobilePhoneNo = customerReq.MobilePhoneNo
	createdCustomer.CustomerName = customerReq.CustomerName
	if err = ac.customerUC.ReadOrCreateCustomer(&createdCustomer); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	createdUser.UserName = customerReq.UserName
	createdUser.UserPassword = customerReq.UserPassword
	createdUser.Email = customerReq.Email
	createdUser.CustomerID = createdCustomer.ID
	createdUser.Encode()
	if err = ac.authUC.CreateUser(&createdUser); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	ac.SuccessResponse(ctx, createdCustomer)
}

func (ac *AuthController) loginUser(ctx *gin.Context) {
	var (
		user     entity.UserCredential
		realUser entity.UserCredential
	)

	err := ac.ParseBodyRequest(ctx, &user)
	if user.UserName == "" {
		ac.FailedResponse(ctx, utils.RequiredError("username"))
		return
	} else if user.UserPassword == "" {
		ac.FailedResponse(ctx, utils.RequiredError("password"))
		return
	} else if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	realUser.UserName = user.UserName
	if err = ac.authUC.FindUser(&realUser); err != nil {
		ac.FailedResponse(ctx, errors.New("wrong username"))
		return
	}
	realUser.Decode()
	if realUser.UserPassword != user.UserPassword {
		ac.FailedResponse(ctx, errors.New("wrong password"))
		return
	}

	generateToken, err := ac.authToken.CreateAccessToken(&user)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	if err := ac.authToken.StoreAccessToken(user.UserName, generateToken); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	ac.SuccessResponse(ctx, generateToken.AccessToken)
}
