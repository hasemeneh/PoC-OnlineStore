package order

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-OnlineStore/svc/order/src/models"
	"github.com/jmoiron/sqlx"
)

var insertOrderQuery = "INSERT INTO `o_orders` (`user_id`,`cart_id`) VALUES (?,?);"

func (p *orderDomain) InsertOrder(ctx context.Context, dbtx *sqlx.Tx, req models.OrderModel) (int64, error) {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	result, err := execFunc(ctx, insertOrderQuery, req.UserID, req.CartID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()

}

func (p *orderDomain) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	return p.DB.BeginTxx(ctx, nil)
}

var getOrderByUserIDQuery = "SELECT id,user_id FROM o_orders WHERE user_id = ?"

func (p *orderDomain) GetOrderByUserID(ctx context.Context, UserID int64) ([]models.OrderModel, error) {
	resp := make([]models.OrderModel, 0)
	err := p.DB.SelectContext(ctx, &resp, getOrderByUserIDQuery, UserID)
	return resp, err

}
