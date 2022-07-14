package tool

import (
	"log"

	"wmb-rest-api/manager"
	"wmb-rest-api/model/entity"
)

func RunSeeds(dbc manager.InfraManagerInterface) {
	var err error

	sqlDB, _ := dbc.DBCon().DB()
	defer sqlDB.Close()

	repoMng := manager.NewRepo(dbc)

	customers := []*entity.Customer{
		{
			CustomerName:  "Kadir",
			MobilePhoneNo: "0877123333",
		},
		{
			CustomerName:  "Devi",
			MobilePhoneNo: "0877745983",
			IsMember:      true,
			Discounts: []*entity.Discount{
				{
					Description: "Diskon Member 10%",
					Pct:         10,
				},
			},
		},
	}
	for _, data := range customers {
		if err = repoMng.CustomerRepo().Create(data); err != nil {
			log.Fatal(err)
		}
	}

	menus := []*entity.Menu{
		{
			MenuName: "Nasi Putih",
			MenuPrices: []entity.MenuPrice{
				{
					Price: 3000,
				},
			},
		},
		{
			MenuName: "Sayur Sop",
			MenuPrices: []entity.MenuPrice{
				{
					Price: 2000,
				},
			},
		},
		{
			MenuName: "Tahu",
			MenuPrices: []entity.MenuPrice{
				{
					Price: 2000,
				},
			},
		},
		{
			MenuName: "Es Teh Tawar",
			MenuPrices: []entity.MenuPrice{
				{
					Price: 1500,
				},
			},
		},
	}
	for _, data := range menus {
		if err = repoMng.MenuRepo().Create(data); err != nil {
			log.Fatal(err)
		}
	}

	tables := []*entity.Table{
		{
			TableDescription: "Table 1",
			IsAvailable:      true,
		},
		{
			TableDescription: "Table 2",
			IsAvailable:      true,
		},
		{
			TableDescription: "Table 3",
			IsAvailable:      true,
		},
	}
	for _, data := range tables {
		if err = repoMng.TableRepo().Create(data); err != nil {
			log.Fatal(err)
		}
	}

	trxTypes := []*entity.TransactionType{
		{
			ID:          "DI",
			Description: "Dine In",
		},
		{
			ID:          "TA",
			Description: "Take Away",
		},
	}
	for _, data := range trxTypes {
		if err = repoMng.TrxTypeRepo().Create(data); err != nil {
			log.Fatal(err)
		}
	}

	updateMenu := menus[0]
	updateMenu.MenuPrices = append(updateMenu.MenuPrices, entity.MenuPrice{Price: 5000})
	if err = repoMng.MenuRepo().Update(updateMenu); err != nil {
		log.Fatal(err)
	}
}
