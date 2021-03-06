// Code generated by mockery v1.0.0. DO NOT EDIT.

package repositories

import (
	context "context"

	models "github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/models"
	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"
)

// MockKeepingRepo is an autogenerated mock type for the KeepingRepo type
type MockKeepingRepo struct {
	mock.Mock
}

// GetProductByProductID provides a mock function with given fields: ctx, productID
func (_m *MockKeepingRepo) GetProductByProductID(ctx context.Context, productID int64) (*models.KeepingModel, error) {
	ret := _m.Called(ctx, productID)

	var r0 *models.KeepingModel
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.KeepingModel); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.KeepingModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByProductIDWithLock provides a mock function with given fields: ctx, dbtx, productID
func (_m *MockKeepingRepo) GetProductByProductIDWithLock(ctx context.Context, dbtx *sqlx.Tx, productID int64) (*models.KeepingModel, error) {
	ret := _m.Called(ctx, dbtx, productID)

	var r0 *models.KeepingModel
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, int64) *models.KeepingModel); ok {
		r0 = rf(ctx, dbtx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.KeepingModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.Tx, int64) error); ok {
		r1 = rf(ctx, dbtx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartTx provides a mock function with given fields: ctx
func (_m *MockKeepingRepo) StartTx(ctx context.Context) (*sqlx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 *sqlx.Tx
	if rf, ok := ret.Get(0).(func(context.Context) *sqlx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStockQuantityByProductID provides a mock function with given fields: ctx, dbtx, productID, quantity
func (_m *MockKeepingRepo) UpdateStockQuantityByProductID(ctx context.Context, dbtx *sqlx.Tx, productID int64, quantity int64) error {
	ret := _m.Called(ctx, dbtx, productID, quantity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, int64, int64) error); ok {
		r0 = rf(ctx, dbtx, productID, quantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
