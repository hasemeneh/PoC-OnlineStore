# PoC-OnlineStore


## Development
- Test User ID integer
- Test Product ID : 101
### Run service
- Run `make docker-start` to start docker container
- Run `make docker-rebuilddb-warehouse` to create the database for warehouse service
- Run `make docker-rebuilddb-order` to create the database for order service
- Run `make docker-rebuilddb-cart` to create the database for cart service
- Open New Terminal Run `make docker-run-warehouse` to start warehouse service
- Open New Terminal Run `make docker-run-order` to start order service
- Open New Terminal Run `make docker-run-cart` to start cart service


For Testing you can use  [this postman](https://github.com/hasemeneh/PoC-OnlineStore/blob/main/files/documents/PoCStore.postman_collection.json)
to test Grpc I recommend you to use [grpcox](https://github.com/gusaul/grpcox)

### Reset Data

- Run `make docker-rebuilddb-warehouse` to create the database for warehouse service
- Run `make docker-rebuilddb-order` to create the database for order service
- Run `make docker-rebuilddb-cart` to create the database for cart service
