package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	TransactionDate   time.Time `gorm:"not null"`
	TableID           sql.NullInt64
	TransactionTypeID string `gorm:"not null"`
	CustomerID        uint   `gorm:"not null"`
	DiscountID        sql.NullInt64

	MenuPrices  []*MenuPrice `gorm:"many2many:t_bill_detail"`
	Discount    Discount
	BillPayment BillPayment
}

func (b Bill) TableName() string {
	return "t_bill"
}

// func (b Bill) String() string {
// 	json, err := json.MarshalIndent(b, "", "  ")
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return string(json)
// }
