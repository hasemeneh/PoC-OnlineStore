package item

import (
	"context"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-OnlineStore/svc/cart/src/models"
	"github.com/jmoiron/sqlx"
)

var insertItemQuery = "INSERT INTO `item` (`product_id`, `quantity`, `cart_id`) VALUES ( ?, ?, ?);"

func (p *itemDomain) InsertItem(ctx context.Context, dbtx *sqlx.Tx, req models.ItemModels) (int64, error) {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	result, err := execFunc(ctx, insertItemQuery, req.ProductID, req.Quantity, req.CartID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()

}

func (p *itemDomain) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	return p.DB.BeginTxx(ctx, nil)
}

var getItemByUserIDQuery = "SELECT i.id,i.cart_id,i.product_id,i.quantity FROM item i INNER JOIN cart c ON c.id = i.cart_id WHERE o.user_id = ?"

func (p *itemDomain) GetItemByUserID(ctx context.Context, UserID int64) (map[int64][]models.ItemModels, error) {
	temp := make([]models.ItemModels, 0)
	err := p.DB.SelectContext(ctx, &temp, getItemByUserIDQuery, UserID)
	resp := make(map[int64][]models.ItemModels)
	for _, item := range temp {
		resp[item.CartID] = append(resp[item.CartID], item)
	}
	return resp, err

}
