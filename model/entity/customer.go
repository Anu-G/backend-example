package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName  string `gorm:"size:100;not null"`
	MobilePhoneNo string `gorm:"size:17;unique;not null"`
	IsMember      bool   `gorm:"not null"`

	Discounts      []*Discount `gorm:"many2many:m_customer_discount"`
	Bills          []Bill
	UserCredential UserCredential
}

func (c Customer) TableName() string {
	return "m_customer"
}

// func (c Customer) String() string {
// 	json, err := json.MarshalIndent(c, "", "  ")
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return string(json)
// }
