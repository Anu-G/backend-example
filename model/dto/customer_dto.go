package dto

type CustomerRequest struct {
	CustomerID    uint   `json:"customer_id"`
	CustomerName  string `json:"customer_name"`
	MobilePhoneNo string `json:"mobile_phone_no"`
	IsMember      bool   `json:"is_member"`
	DiscountID    uint   `json:"discount_id"`
}

type RegisterCustomerRequest struct {
	UserName      string `json:"user_name"`
	UserPassword  string `json:"user_password"`
	CustomerName  string `json:"customer_name"`
	MobilePhoneNo string `json:"mobile_phone_no"`
	Email         string `json:"email"`
}
