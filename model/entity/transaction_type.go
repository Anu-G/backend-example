package entity

type TransactionType struct {
	ID          string `gorm:"primaryKey"`
	Description string

	Bills []Bill
}

func (tt TransactionType) TableName() string {
	return "m_trans_type"
}

// func (tt TransactionType) String() string {
// 	json, err := json.MarshalIndent(tt, "", "  ")
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return string(json)
// }
