pkgs          = $(shell go list ./... | grep -vE '(vendor|mock)')
NOW=$(shell date)
compose_file=./docker/docker-compose.yml
compose=docker-compose -f ${compose_file}
cart_service_binary=cart
order_service_binary=order
warehouse_service_binary=warehouse
current_dir=$(shell pwd)


docker-start:
	@echo "${NOW} BUILDING SAMPLE..."
	@echo "${NOW} STARTING CONTAINER..."
	@${compose} up -d --build

docker-stop:
	@echo "${NOW} STOPPING CONTAINER..."
	@${compose} down
	@echo "${NOW} CLEAN UP..."
	@rm -f ./bin/cart/${cart_service_binary} 
	@rm -f ./bin/order/${order_service_binary} 
	@rm -f ./bin/warehouse/${warehouse_service_binary} 

docker-rebuilddb:
	@echo "${NOW} REBUILDDB..."
	@echo "${NOW} DROPING EXISTING DB..."
	@docker exec -it wi-db  mysql -uroot -e'drop database if exists poc_store'
	@echo "${NOW} CREATE DB..."
	@docker exec -it wi-db  mysql -uroot -e'create database poc_store'
	@echo "${NOW} RUN SQL SCRIPTS..."
	@docker exec -it wi-db setup.sh /etc/database

docker-run-cart:
	@echo "${NOW} BUILDING..."
	@cd ./svc/cart/src && go build -o ./../../../bin/cart/${cart_service_binary}
	@echo "${NOW} RUNNING..."
	@docker exec -it cart /usr/local/bin/${cart_service_binary}

docker-run-order:
	@echo "${NOW} BUILDING..."
	@cd ./svc/order/src && go build -o ./../../../bin/order/${order_service_binary}
	@echo "${NOW} RUNNING..."
	@docker exec -it order /usr/local/bin/${order_service_binary}

docker-run-warehouse:
	@echo "${NOW} BUILDING..."
	@cd ./svc/warehouse/src && go build -o ./../../../bin/warehouse/${warehouse_service_binary}
	@echo "${NOW} RUNNING..."
	@docker exec -it warehouse /usr/local/bin/${warehouse_service_binary}

docker-run:
	@echo "${NOW} BUILDING..."
	@cd ./svc/cart/src && go build -o ./../../../bin/cart/${cart_service_binary}
	@cd ./svc/order/src && go build -o ./../../../bin/order/${order_service_binary}
	@cd ./svc/warehouse/src && go build -o ./../../../bin/warehouse/${warehouse_service_binary}
	@echo "${NOW} RUNNING..."
	@docker exec -it warehouse /usr/local/bin/${warehouse_service_binary}
	@docker exec -it cart /usr/local/bin/${cart_service_binary}
	@docker exec -it order /usr/local/bin/${order_service_binary}
