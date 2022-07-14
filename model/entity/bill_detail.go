package entity

type BillDetail struct {
	ID          uint `gorm:"primaryKey"`
	BillID      uint `gorm:"primaryKey;not null"`
	MenuPriceID uint `gorm:"primaryKey;not null"`
	Qty         uint `gorm:"default:0;not null"`
}

func (bd BillDetail) TableName() string {
	return "t_bill_detail"
}

// func (bd BillDetail) String() string {
// 	json, err := json.MarshalIndent(bd, "", "  ")
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return string(json)
// }
