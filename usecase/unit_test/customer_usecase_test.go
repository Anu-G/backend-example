package unit_test

import (
	"testing"

	"wmb-rest-api/model/entity"
	rmock "wmb-rest-api/repository/mocks"
	"wmb-rest-api/usecase"
	umock "wmb-rest-api/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

var dummyCustomer = []entity.Customer{
	{
		Model:         gorm.Model{ID: 1},
		CustomerName:  "Andi",
		MobilePhoneNo: "0813",
	},
	{
		Model:         gorm.Model{ID: 2},
		CustomerName:  "Budi",
		MobilePhoneNo: "0812",
	},
}

type CustomerUCTest struct {
	suite.Suite
	customerMock *rmock.CustomerRepositoryInterface
	discountMock *umock.DiscountUseCaseInterface
}

func (suite *CustomerUCTest) SetupTest() {
	suite.customerMock = rmock.NewCustomerRepositoryInterface(suite.T())
	suite.discountMock = umock.NewDiscountUseCaseInterface(suite.T())
}

func Test_CustomerUCRun(t *testing.T) {
	suite.Run(t, new(CustomerUCTest))
}

func (suite *CustomerUCTest) Test_ReadOrCreateCustomer_Pos() {
	// dummy := dummyCustomer[0]
	req := entity.Customer{
		MobilePhoneNo: "0813",
	}

	mapSearch := map[string]interface{}{"mobile_phone_no": req.MobilePhoneNo}
	suite.customerMock.On("FindFirtst", mapSearch).Return(req, nil)
	// suite.customerMock.On("Create", &req).Return(nil).Run(func(args mock.Arguments) {
	// 	arg := args.Get(0).(*entity.Customer)
	// 	log.Println(arg)
	// 	arg.Model = dummy.Model
	// 	arg.CustomerName = dummy.CustomerName
	// })

	customerUCTest := usecase.NewCustomerUseCase(suite.customerMock, suite.discountMock)

	err := customerUCTest.ReadOrCreateCustomer(&req)
	assert.Nil(suite.T(), err)
	// assert.Equal(suite.T(), dummy, req)
}
