package models

type MainConfig struct {
	DBConnectionString string `json:"db_connection"`
	RunningPort        string `json:"running_port"`
	WarehouseAddress   string `json:"warehouse_address"`
	GRPCPort           string `json:"grpc_port"`
}
