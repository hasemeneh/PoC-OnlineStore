package models

type ItemModels struct {
	ID        int64 `json:"id" db:"id"`
	OrderID   int64 `json:"order_id" db:"order_id"`
	ProductID int64 `json:"product_id" db:"product_id"`
	Quantity  int64 `json:"quantity" db:"quantity"`
}
