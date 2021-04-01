package models

type StockResponse struct {
	*KeepingModel
	History []StockHistoryModel `json:"history"`
}
