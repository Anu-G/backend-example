package tool

import (
	"wmb-rest-api/manager"
	"wmb-rest-api/model/entity"
)

func RunMigrate(dbc manager.InfraManagerInterface) error {
	var err error

	sqlDB, _ := dbc.DBCon().DB()
	defer sqlDB.Close()

	err = dbc.DBCon().SetupJoinTable(&entity.Bill{}, "MenuPrices", &entity.BillDetail{})
	if err != nil {
		panic(err)
	}

	err = dbc.DBCon().SetupJoinTable(&entity.MenuPrice{}, "Bills", &entity.BillDetail{})
	if err != nil {
		panic(err)
	}

	err = dbc.DBCon().AutoMigrate(
		&entity.Menu{}, &entity.Table{}, &entity.TransactionType{}, &entity.Customer{},
		&entity.Discount{}, &entity.MenuPrice{}, &entity.Bill{}, &entity.BillPayment{},
		&entity.UserCredential{})
	if err != nil {
		panic(err)
	}
	return err
}
