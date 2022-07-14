package controller

import (
	"fmt"

	"wmb-rest-api/delivery/api"
	"wmb-rest-api/model/dto"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/usecase"
	"wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	router  *gin.Engine
	usecase usecase.TrxUseCaseInterface
	api.BaseApi
}

func NewTransactionController(router *gin.Engine, uc usecase.TrxUseCaseInterface) *TransactionController {
	controller := TransactionController{
		router:  router,
		usecase: uc,
	}

	routeTransaction := controller.router.Group("/transaction")
	routeTransaction.POST("/create", controller.createTransaction)
	// routeTransaction.GET("/print", controller.printTransaction)
	routeTransaction.GET("/revenue", controller.dailyRevenue)
	routeTransaction.POST("/payment/balance", controller.checkBalance)
	routeTransaction.POST("/payment/pay", controller.doPayment)

	return &controller
}

func (tc *TransactionController) createTransaction(ctx *gin.Context) {
	var (
		transactionReq dto.CreateTransaction
		billID         int
	)

	err := tc.ParseBodyRequest(ctx, &transactionReq)
	if transactionReq.Customer.MobilePhoneNo == "" {
		tc.FailedResponse(ctx, utils.RequiredError("mobile phone number"))
		return
	} else if transactionReq.TransactionTypeID == "" {
		tc.FailedResponse(ctx, utils.RequiredError("transaction type id"))
		return
	} else if transactionReq.OrderMenus == nil {
		tc.FailedResponse(ctx, utils.RequiredError("order"))
		return
	} else if err != nil {
		tc.FailedResponse(ctx, err)
		return
	}

	if billID, err = tc.usecase.CreateTransaction(&transactionReq); err != nil {
		tc.FailedResponse(ctx, err)
		return
	}
	resp := fmt.Sprintf("transaction created! id:%v", billID)
	tc.SuccessResponse(ctx, resp)
}

// func (tc *TransactionController) printTransaction(ctx *gin.Context) {
// 	var trx entity.Bill

// 	err := tc.ParseBodyRequest(ctx, &trx)
// 	if trx.ID == 0 {
// 		tc.FailedResponse(ctx, utils.RequiredError("id"))
// 		return
// 	} else if err != nil {
// 		tc.FailedResponse(ctx, err)
// 		return
// 	}

// 	if printOut, err := tc.usecase.PrintAndFinishTransaction(&trx); err != nil {
// 		tc.FailedResponse(ctx, err)
// 		return
// 	} else {
// 		tc.SuccessResponse(ctx, printOut)
// 	}
// }

func (tc *TransactionController) dailyRevenue(ctx *gin.Context) {
	var rev dto.Revenue

	err := tc.ParseBodyRequest(ctx, &rev)
	if rev.TransactionDate == "" {
		tc.FailedResponse(ctx, utils.RequiredError("date"))
		return
	} else if err != nil {
		tc.FailedResponse(ctx, err)
		return
	}

	if err := tc.usecase.GetRevenue(&rev); err != nil {
		tc.FailedResponse(ctx, err)
		return
	}
	tc.SuccessResponse(ctx, rev)
}

func (tc *TransactionController) checkBalance(ctx *gin.Context) {
	var cus entity.Customer

	err := tc.ParseBodyRequest(ctx, &cus)
	if cus.MobilePhoneNo == "" {
		tc.FailedResponse(ctx, utils.RequiredError("phone number"))
		return
	} else if err != nil {
		tc.FailedResponse(ctx, err)
		return
	}

	balance, err := tc.usecase.CheckBalance(&cus)
	if err != nil {
		tc.FailedResponse(ctx, err)
		return
	}
	tc.SuccessResponse(ctx, balance)
}

func (tc *TransactionController) doPayment(ctx *gin.Context) {
	var pay dto.PaymentMethod

	err := tc.ParseBodyRequest(ctx, &pay)
	if pay.BillId == 0 {
		tc.FailedResponse(ctx, utils.RequiredError("bill id"))
		return
	} else if pay.PaymentMethod == "" {
		tc.FailedResponse(ctx, utils.RequiredError("payment method"))
		return
	} else if err != nil {
		tc.FailedResponse(ctx, err)
		return
	}

	if printOut, err := tc.usecase.PayAndFinishTransaction(&pay); err != nil {
		tc.FailedResponse(ctx, err)
		return
	} else {
		tc.SuccessResponse(ctx, printOut)
	}
}
