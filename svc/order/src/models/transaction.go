package models

import "github.com/hasemeneh/PoC-OnlineStore/svc/order/src/constants"

type OrderCreationRequest struct {
	UserID   int64
	CartID   int64
	Products []ProductRequest
}

type OrderCreationResponse struct {
	IsSuccess bool
	Products  []ProductResponse
}

type ProductRequest struct {
	ProductID int64 `json:"product_id"`
	Demand    int64 `json:"demand"`
}
type ProductResponse struct {
	ProductRequest
	Status constants.ClaimStatus `json:"status"`
}
type ClaimResponse struct {
	Products []ProductResponse `json:"products"`
}

type ClaimRequest struct {
	OrderID  int64            `json:"order_id"`
	CartID   int64            `json:"cart_id"`
	Products []ProductRequest `json:"products"`
}
