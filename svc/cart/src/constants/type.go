package constants

type HistoryType int

const (
	HistoryTypeOutcoming = HistoryType(50)
	HistoryTypeIncoming  = HistoryType(100)
)

type ClaimStatus int

const (
	ClaimStatusOK         = ClaimStatus(50)
	ClaimStatusOutOfStock = ClaimStatus(100)
)
