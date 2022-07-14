package entity

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	TableDescription string
	IsAvailable      bool `gorm:"not null"`

	Bills []Bill
}

func (t Table) TableName() string {
	return "m_table"
}

// func (t Table) String() string {
// 	json, err := json.MarshalIndent(t, "", "  ")
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return string(json)
// }
