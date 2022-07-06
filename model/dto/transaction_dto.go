package dto

import (
	"wmb-rest-api/model/entity"
)

type CreateTransaction struct {
	TableID           uint   `json:"table_id"`
	TransactionTypeID string `json:"transaction_type_id"`

	OrderMenus []OrderMenu     `json:"order_menu"`
	Customer   entity.Customer `json:"customer"`
}

type OrderMenu struct {
	MenuID uint `json:"menu_id"`
	Qty    uint `json:"qty"`
}

type BillPrintOut struct {
	BillID          uint    `json:"bill_id"`
	TransactionDate string  `json:"transaction_date"`
	CustomerName    string  `json:"customer_name"`
	TransactionType string  `json:"transaction_type"`
	Table           string  `json:"table_number,omitempty"`
	Discount        uint    `json:"discount,omitempty"`
	GrandTotal      float64 `json:"grand_total"`

	Orders []HistoryMenuOrder `json:"order_menu"`
}

type HistoryMenuOrder struct {
	MenuName string  `json:"menu_name"`
	Price    float64 `json:"menu_price"`
	Qty      uint    `json:"qty"`
	Subtotal float64 `json:"subtotal"`
}

type Revenue struct {
	TransactionDate string  `json:"transaction_date"`
	TotalRevenue    float64 `json:"total_revenue"`
}
