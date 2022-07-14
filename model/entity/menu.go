package entity

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	MenuName string `gorm:"size:100;unique;not null"`

	MenuPrices []MenuPrice
}

func (m Menu) TableName() string {
	return "m_menu"
}

// func (m Menu) String() string {
// 	json, err := json.MarshalIndent(m, "", "  ")
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return string(json)
// }
