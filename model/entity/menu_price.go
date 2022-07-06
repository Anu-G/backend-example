package entity

import (
	"encoding/json"

	"gorm.io/gorm"
)

type MenuPrice struct {
	gorm.Model
	MenuID uint
	Price  float64 `gorm:"default:0;not null"`

	Bills []*Bill `gorm:"many2many:t_bill_detail"`
}

func (mp MenuPrice) TableName() string {
	return "m_menu_price"
}

func (mp MenuPrice) String() string {
	json, err := json.MarshalIndent(mp, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(json)
}
