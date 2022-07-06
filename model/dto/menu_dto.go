package dto

type MenuRequest struct {
	MenuID    uint    `json:"menu_id"`
	MenuName  string  `json:"menu_name"`
	MenuPrice float64 `json:"menu_price"`
}
