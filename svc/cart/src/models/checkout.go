package models

type CheckoutResponse struct {
	IsSuccess bool              `json:"is_success"`
	Products  []ProductResponse `json:"products"`
}

type UserRecordedCheckout struct {
	CartModel
	Item []ItemModels `json:"items"`
}
