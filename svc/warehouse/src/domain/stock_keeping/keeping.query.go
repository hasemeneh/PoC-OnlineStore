package stock_keeping

import (
	"context"
	"errors"

	"github.com/hasemeneh/PoC-OnlineStore/helper/database"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
	"github.com/jmoiron/sqlx"
)

func (p *stockKeepingDomain) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	return p.DB.BeginTxx(ctx, nil)
}

var getProductByProductIDQuery = "SELECT id,product_id,quantity FROM wh_stock_keeping WHERE product_id = ?"

func (p *stockKeepingDomain) GetProductByProductID(ctx context.Context, productID int64) (*models.KeepingModel, error) {
	resp := models.KeepingModel{}
	err := p.DB.GetContext(ctx, &resp, getProductByProductIDQuery, productID)
	return &resp, err

}

var getProductByProductIDForUpdateQuery = "SELECT id,product_id,quantity FROM wh_stock_keeping WHERE product_id = ? FOR UPDATE"

func (p *stockKeepingDomain) GetProductByProductIDWithLock(ctx context.Context, dbtx *sqlx.Tx, productID int64) (*models.KeepingModel, error) {
	resp := models.KeepingModel{}
	if dbtx == nil {
		return nil, errors.New("DBTX required")
	}
	err := dbtx.GetContext(ctx, &resp, getProductByProductIDForUpdateQuery, productID)
	return &resp, err

}

var updateStockByProductID = "UPDATE `wh_stock_keeping` SET `quantity` = ? WHERE `wh_stock_keeping`.`product_id` = ?;"

func (p *stockKeepingDomain) UpdateStockQuantityByProductID(ctx context.Context, dbtx *sqlx.Tx, productID int64, quantity int64) error {
	var execFunc database.ExecFunc
	if dbtx == nil {
		execFunc = p.DB.ExecContext
	} else {
		execFunc = dbtx.ExecContext
	}
	_, err := execFunc(ctx, updateStockByProductID, quantity, productID)
	return err

}
