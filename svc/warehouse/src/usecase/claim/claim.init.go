package claim

import "github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/repositories"

type claimUsecase struct {
	HistoryRepo repositories.HistoryRepo
	KeepingRepo repositories.KeepingRepo
}
type Option struct {
	HistoryRepo repositories.HistoryRepo
	KeepingRepo repositories.KeepingRepo
}

func New(o *Option) *claimUsecase {
	return &claimUsecase{
		HistoryRepo: o.HistoryRepo,
		KeepingRepo: o.KeepingRepo,
	}
}
