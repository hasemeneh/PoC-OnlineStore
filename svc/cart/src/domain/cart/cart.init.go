package cart

import (
	"github.com/hasemeneh/PoC-OnlineStore/helper/redis"
	"github.com/jmoiron/sqlx"
)

type cartDomain struct {
	DB    *sqlx.DB
	redis redis.RedisInterface
}

func New(DB *sqlx.DB, redis redis.RedisInterface) *cartDomain {
	return &cartDomain{
		DB:    DB,
		redis: redis,
	}
}
