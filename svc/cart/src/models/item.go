package models

type ItemModels struct {
	ID        int64 `json:"id" db:"id"`
	CartID    int64 `json:"cart_id" db:"cart_id"`
	ProductID int64 `json:"product_id" db:"product_id"`
	Quantity  int64 `json:"quantity" db:"quantity"`
}
