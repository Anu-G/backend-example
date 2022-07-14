package dto

type LopeiBalance struct {
	CustomerName  string
	MobilePhoneNo string
	Balance       float64
}

type PaymentMethod struct {
	BillId        uint
	PaymentMethod string
}
