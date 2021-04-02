package models

type UserOrder struct {
	OrderModel
	Items []ItemModels `json:"items"`
}
