package models

type AddToCartResponse struct {
	Products []ProductRequest `json:"products"`
}
type AddToCartRequest struct {
	ProductRequest
	UserID int64 `json:"user_id"`
}
