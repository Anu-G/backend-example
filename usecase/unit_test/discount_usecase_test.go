package unit_test

import (
	"testing"

	"wmb-rest-api/model/entity"
	"wmb-rest-api/repository/mocks"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

var dummyDiscount = []entity.Discount{
	{
		Model: gorm.Model{ID: 1},
		Pct:   10,
	},
	{
		Model: gorm.Model{ID: 2},
		Pct:   15,
	},
}

type DiscountUCTest struct {
	suite.Suite
	discountMock *mocks.DiscountRepositoryInterface
}

func (suite *DiscountUCTest) SetupTest() {
	suite.discountMock = mocks.NewDiscountRepositoryInterface(suite.T())
}

func Test_DiscountUCRun(t *testing.T) {
	suite.Run(t, new(DiscountUCTest))
}
