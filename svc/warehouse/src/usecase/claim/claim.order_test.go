package claim

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/constants"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/repositories"
)

func TestValidateProductStockSuccess(t *testing.T) {
	mockHistory := repositories.MockHistoryRepo{}
	mockKeeping := repositories.MockKeepingRepo{}
	module := New(&Option{
		HistoryRepo: &mockHistory,
		KeepingRepo: &mockKeeping,
	})
	g := NewGomegaWithT(t)
	mockKeeping.On("GetProductByProductID", mock.Anything, mock.Anything).Return(&models.KeepingModel{
		Quantity: 10000,
	}, nil)
	resp, err := module.ValidateProductStock(context.Background(), &models.ProductRequest{
		Demand: 1,
	})
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(resp.Status).Should(Equal(constants.ClaimStatusOK))
}

func TestValidateProductStockFailedGetProduct(t *testing.T) {
	mockHistory := repositories.MockHistoryRepo{}
	mockKeeping := repositories.MockKeepingRepo{}
	module := New(&Option{
		HistoryRepo: &mockHistory,
		KeepingRepo: &mockKeeping,
	})
	g := NewGomegaWithT(t)
	mockKeeping.On("GetProductByProductID", mock.Anything, mock.Anything).Return(&models.KeepingModel{
		Quantity: 10000,
	}, errors.New("error"))
	resp, err := module.ValidateProductStock(context.Background(), &models.ProductRequest{
		Demand: 1,
	})
	g.Expect(err).Should(HaveOccurred())
	g.Expect(resp.Status).ShouldNot(Equal(constants.ClaimStatusOK))
}

func TestValidateProductStockBadReq(t *testing.T) {
	mockHistory := repositories.MockHistoryRepo{}
	mockKeeping := repositories.MockKeepingRepo{}
	module := New(&Option{
		HistoryRepo: &mockHistory,
		KeepingRepo: &mockKeeping,
	})
	g := NewGomegaWithT(t)

	resp, err := module.ValidateProductStock(context.Background(), nil)
	g.Expect(err).Should(HaveOccurred())
	g.Expect(resp.Status).ShouldNot(Equal(constants.ClaimStatusOK))
}

func TestClaimProductSuccess(t *testing.T) {
	g := NewGomegaWithT(t)
	mockHistory := repositories.MockHistoryRepo{}
	mockKeeping := repositories.MockKeepingRepo{}
	db, dbMock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dbMock.ExpectBegin()
	dbtx, _ := sqlxDB.BeginTxx(context.Background(), nil)
	dbMock.ExpectCommit()
	module := New(&Option{
		HistoryRepo: &mockHistory,
		KeepingRepo: &mockKeeping,
	})
	mockKeeping.On("GetProductByProductIDWithLock", mock.Anything, mock.Anything, mock.Anything).Return(&models.KeepingModel{}, nil)
	mockKeeping.On("UpdateStockQuantityByProductID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockKeeping.On("StartTx", mock.Anything).Return(dbtx, nil)
	mockHistory.On("InsertHistory", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	_, err := module.ClaimProduct(context.Background(), &models.ClaimRequest{
		Products: []models.ProductRequest{
			{},
			{},
			{},
		},
	})
	g.Expect(err).ShouldNot(HaveOccurred())
}

func TestClaimProductErrInserHistory(t *testing.T) {
	g := NewGomegaWithT(t)
	mockHistory := repositories.MockHistoryRepo{}
	mockKeeping := repositories.MockKeepingRepo{}
	db, dbMock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dbMock.ExpectBegin()
	dbtx, _ := sqlxDB.BeginTxx(context.Background(), nil)
	dbMock.ExpectCommit()
	module := New(&Option{
		HistoryRepo: &mockHistory,
		KeepingRepo: &mockKeeping,
	})
	mockKeeping.On("GetProductByProductIDWithLock", mock.Anything, mock.Anything, mock.Anything).Return(&models.KeepingModel{}, nil)
	mockKeeping.On("UpdateStockQuantityByProductID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockKeeping.On("StartTx", mock.Anything).Return(dbtx, nil)
	mockHistory.On("InsertHistory", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))

	_, err := module.ClaimProduct(context.Background(), &models.ClaimRequest{
		Products: []models.ProductRequest{
			{},
		},
	})
	g.Expect(err).Should(HaveOccurred())
}

func TestClaimProductErrUpdateStockQuantityByProductID(t *testing.T) {
	g := NewGomegaWithT(t)
	mockHistory := repositories.MockHistoryRepo{}
	mockKeeping := repositories.MockKeepingRepo{}
	db, dbMock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dbMock.ExpectBegin()
	dbtx, _ := sqlxDB.BeginTxx(context.Background(), nil)
	dbMock.ExpectCommit()
	module := New(&Option{
		HistoryRepo: &mockHistory,
		KeepingRepo: &mockKeeping,
	})
	mockKeeping.On("StartTx", mock.Anything).Return(dbtx, nil)
	mockKeeping.On("GetProductByProductIDWithLock", mock.Anything, mock.Anything, mock.Anything).Return(&models.KeepingModel{}, nil)
	mockKeeping.On("UpdateStockQuantityByProductID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))

	_, err := module.ClaimProduct(context.Background(), &models.ClaimRequest{
		Products: []models.ProductRequest{
			{},
		},
	})
	g.Expect(err).Should(HaveOccurred())
}

func TestClaimProductErrGetProductByProductIDWithLock(t *testing.T) {
	g := NewGomegaWithT(t)
	mockHistory := repositories.MockHistoryRepo{}
	mockKeeping := repositories.MockKeepingRepo{}
	db, dbMock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dbMock.ExpectBegin()
	dbtx, _ := sqlxDB.BeginTxx(context.Background(), nil)
	dbMock.ExpectCommit()
	module := New(&Option{
		HistoryRepo: &mockHistory,
		KeepingRepo: &mockKeeping,
	})
	mockKeeping.On("StartTx", mock.Anything).Return(dbtx, nil)
	mockKeeping.On("GetProductByProductIDWithLock", mock.Anything, mock.Anything, mock.Anything).Return(&models.KeepingModel{}, errors.New("error"))

	_, err := module.ClaimProduct(context.Background(), &models.ClaimRequest{
		Products: []models.ProductRequest{
			{},
		},
	})
	g.Expect(err).Should(HaveOccurred())
}

func TestClaimProductErrStartTx(t *testing.T) {
	g := NewGomegaWithT(t)
	mockHistory := repositories.MockHistoryRepo{}
	mockKeeping := repositories.MockKeepingRepo{}
	db, dbMock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dbMock.ExpectBegin()
	dbtx, _ := sqlxDB.BeginTxx(context.Background(), nil)
	dbMock.ExpectCommit()
	module := New(&Option{
		HistoryRepo: &mockHistory,
		KeepingRepo: &mockKeeping,
	})
	mockKeeping.On("StartTx", mock.Anything).Return(dbtx, errors.New("error"))

	_, err := module.ClaimProduct(context.Background(), &models.ClaimRequest{
		Products: []models.ProductRequest{
			{},
		},
	})
	g.Expect(err).Should(HaveOccurred())
}
