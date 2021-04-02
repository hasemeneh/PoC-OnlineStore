package order

import (
	"github.com/jmoiron/sqlx"
)

type orderDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *orderDomain {
	return &orderDomain{
		DB: DB,
	}
}
