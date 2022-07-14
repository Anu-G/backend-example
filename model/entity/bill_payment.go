package entity

import (
	"gorm.io/gorm"
)

type BillPayment struct {
	gorm.Model
	BillID        uint    `gorm:"not null"`
	PaymentMethod string  `gorm:"not null"`
	TotalPayment  float64 `gorm:"not null"`
}

func (bp BillPayment) TableName() string {
	return "t_bill_payment"
}

// func (bp BillPayment) String() string {
// 	json, err := json.MarshalIndent(bp, "", "  ")
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return string(json)
// }
