package entity

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName  string `gorm:"size:100;not null"`
	MobilePhoneNo string `gorm:"size:17;unique;not null"`
	IsMember      bool   `gorm:"default:false;not null"`

	Discounts []*Discount `gorm:"many2many:m_customer_discount"`
	Bills     []Bill
}

func (c Customer) TableName() string {
	return "m_customer"
}

func (c Customer) String() string {
	json, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(json)
}
