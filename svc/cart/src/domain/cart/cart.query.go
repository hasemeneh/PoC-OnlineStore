package cart

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	"github.com/jmoiron/sqlx"
)

var insertCartQuery = "INSERT INTO `cart` (`user_id`) VALUES (?);"

func (p *cartDomain) InsertCart(ctx context.Context, dbtx *sqlx.Tx, req models.CartModel) (int64, error) {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	result, err := execFunc(ctx, insertCartQuery, req.UserID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()

}

func (p *cartDomain) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	return p.DB.BeginTxx(ctx, nil)
}

var getCartByUserIDQuery = "SELECT id,user_id FROM cart WHERE user_id = ?"

func (p *cartDomain) GetRecordedCartByUserID(ctx context.Context, UserID int64) ([]models.CartModel, error) {
	resp := make([]models.CartModel, 0)
	err := p.DB.SelectContext(ctx, &resp, getCartByUserIDQuery, UserID)
	return resp, err

}
