package stock_keeping

import (
	"github.com/jmoiron/sqlx"
)

type stockKeepingDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *stockKeepingDomain {
	return &stockKeepingDomain{
		DB: DB,
	}
}
