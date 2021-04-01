package stock_history

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
	"github.com/jmoiron/sqlx"
)

func (p *stockHistoryDomain) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	return p.DB.BeginTxx(ctx, nil)
}

var getProductByProductIDQuery = "SELECT h.id,h.type,h.amount,h.additional_info, h.created_at FROM wh_stock_history h INNER JOIN wh_stock_keeping k ON k.id = h.sk_id WHERE k.product_id = ?"

func (p *stockHistoryDomain) GetProductByProductID(ctx context.Context, productID int64) ([]models.StockHistoryModel, error) {
	resp := make([]models.StockHistoryModel, 0)
	err := p.DB.SelectContext(ctx, &resp, getProductByProductIDQuery, productID)
	return resp, err

}

var insertHistoryQuery = "INSERT INTO `wh_stock_history` (`sk_id`, `type`, `amount`, `additional_info`) VALUES ( ?, ?, ?, ?);"

func (p *stockHistoryDomain) InsertHistory(ctx context.Context, dbtx *sqlx.Tx, productID int64, quantity int) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, insertHistoryQuery, quantity, productID)
	return err

}
