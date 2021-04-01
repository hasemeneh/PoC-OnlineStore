package guest

import (
	"github.com/jmoiron/sqlx"
)

type guestDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *guestDomain {
	return &guestDomain{
		DB: DB,
	}
}
