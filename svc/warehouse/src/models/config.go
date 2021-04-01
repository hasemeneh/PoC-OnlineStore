package models

type MainConfig struct {
	DBConnectionString string `json:"db_connection"`
	RunningPort        string `json:"running_port"`
}
