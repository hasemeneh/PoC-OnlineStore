package stock_history

import (
	"github.com/jmoiron/sqlx"
)

type stockHistoryDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *stockHistoryDomain {
	return &stockHistoryDomain{
		DB: DB,
	}
}
