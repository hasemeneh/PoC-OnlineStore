package service

import (
	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/definitions"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/domain/stock_history"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/domain/stock_keeping"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/usecase/claim"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/usecase/control"
)

type Service struct {
	cfg     *models.MainConfig
	Usecase struct {
		Control definitions.ControlDefinition
		Claim   definitions.ClaimDefinition
	}
}

func New(cfg *models.MainConfig) *Service {
	db := database.New().Connect(cfg.DBConnectionString)
	historyDomain := stock_history.New(db)
	keepingDomain := stock_keeping.New(db)

	serviceObj := Service{
		cfg: cfg,
	}
	serviceObj.Usecase.Claim = claim.New(&claim.Option{
		HistoryRepo: historyDomain,
		KeepingRepo: keepingDomain,
	})
	serviceObj.Usecase.Control = control.New(&control.Option{
		HistoryRepo: historyDomain,
		KeepingRepo: keepingDomain,
	})
	return &serviceObj
}
