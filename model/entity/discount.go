package entity

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	Description string
	Pct         uint `gorm:"default:0;not null"`

	Customers []*Customer `gorm:"many2many:m_customer_discount"`
}

func (d Discount) TableName() string {
	return "m_discount"
}

func (d Discount) String() string {
	json, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(json)
}
