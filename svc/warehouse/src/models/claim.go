package models

import "github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/constants"

type ClaimResponse struct {
	Products []ProductResponse `json:"products"`
}
type ProductResponse struct {
	*ProductRequest
	Status constants.ClaimStatus `json:"status"`
}
type ClaimRequest struct {
	Products []ProductRequest `json:"products"`
}
type ProductRequest struct {
	ProductID int64 `json:"product_id"`
	Demand    int64 `json:"demand"`
}
