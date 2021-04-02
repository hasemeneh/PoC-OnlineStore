package models

type MainConfig struct {
	DBConnectionString string `json:"db_connection"`
	RunningPort        string `json:"running_port"`
	RedisConnection    string `json:"redis_connection"`
	WarehouseAddress   string `json:"warehouse_address"`
	OrderAddress       string `json:"order_address"`
	GRPCPort           string `json:"grpc_port"`
}
