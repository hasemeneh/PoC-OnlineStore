pkgs          = $(shell go list ./... | grep -vE '(vendor|mock)')
NOW=$(shell date)
compose_file=./docker/docker-compose.yml
compose=docker-compose -f ${compose_file}
service_binary=poc_store
current_dir=$(shell pwd)


docker-start:
	@echo "${NOW} BUILDING SAMPLE..."
	@echo "${NOW} STARTING CONTAINER..."
	@${compose} up -d --build

docker-stop:
	@echo "${NOW} STOPPING CONTAINER..."
	@${compose} down
	@echo "${NOW} CLEAN UP..."
	@rm -f ./bin/${service_binary} 

docker-rebuilddb:
	@echo "${NOW} REBUILDDB..."
	@echo "${NOW} DROPING EXISTING DB..."
	@docker exec -it wi-db  mysql -uroot -e'drop database if exists poc_store'
	@echo "${NOW} CREATE DB..."
	@docker exec -it wi-db  mysql -uroot -e'create database poc_store'
	@echo "${NOW} RUN SQL SCRIPTS..."
	@docker exec -it wi-db setup.sh /etc/database

docker-run:
	@echo "${NOW} BUILDING..."
	@cd ./src && go build -o ./../bin/${service_binary}
	@echo "${NOW} RUNNING..."
	@docker exec -it poc /usr/local/bin/${service_binary}
