package models

import (
	"time"

	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/constants"
)

type StockHistoryModel struct {
	ID             int64                 `json:"id" db:"id"`
	Type           constants.HistoryType `json:"type" db:"type"`
	Amount         int64                 `json:"amount" db:"amount"`
	AdditionalInfo MetaData              `json:"additional_info" db:"additional_info"`
	CreatedAt      time.Time             `json:"created_at" db:"created_at"`
}
