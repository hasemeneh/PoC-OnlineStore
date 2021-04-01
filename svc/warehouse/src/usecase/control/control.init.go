package control

import "github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/repositories"

type controlUsecase struct {
	HistoryRepo repositories.HistoryRepo
	KeepingRepo repositories.KeepingRepo
}
type Option struct {
	HistoryRepo repositories.HistoryRepo
	KeepingRepo repositories.KeepingRepo
}

func New(o *Option) *controlUsecase {
	return &controlUsecase{
		HistoryRepo: o.HistoryRepo,
		KeepingRepo: o.KeepingRepo,
	}
}
