module github.com/hasemeneh/PoC-OnlineStore/svc/cart

go 1.14

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/hasemeneh/PoC-OnlineStore v0.0.0-20210402080515-f297e0df476e
	github.com/hasemeneh/PoC-OnlineStore/helper v0.0.0-20210402022656-782bf15ce525
	github.com/hasemeneh/PoC-OnlineStore/svc/order v0.0.0-20210402022656-782bf15ce525 // indirect
	github.com/jmoiron/sqlx v1.3.1
	github.com/julienschmidt/httprouter v1.3.0
	google.golang.org/grpc v1.36.1
)
