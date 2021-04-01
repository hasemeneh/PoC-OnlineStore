package service

import "github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"

type Service struct {
	cfg *models.MainConfig
}

func New(cfg *models.MainConfig) *Service {
	// db := database.New().Connect(cfg.DBConnectionString)
	serviceObj := Service{}
	return &serviceObj
}
