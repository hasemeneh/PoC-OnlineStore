package models

type OrderModel struct {
	ID     int64 `json:"id" db:"id"`
	UserID int64 `json:"user_id" db:"user_id"`
	CartID int64 `json:"cart_id" db:"cart_id"`
}
