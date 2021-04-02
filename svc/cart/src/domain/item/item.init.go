package item

import (
	"github.com/jmoiron/sqlx"
)

type itemDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *itemDomain {
	return &itemDomain{
		DB: DB,
	}
}
