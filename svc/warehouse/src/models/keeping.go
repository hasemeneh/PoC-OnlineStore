package models

type KeepingModel struct {
	ID        int64 `json:"id" db:"id"`
	ProductID int64 `json:"product_id" db:"product_id"`
	Quantity  int64 `json:"quantity" db:"quantity"`
}
