// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/kumin/BityDating/entities"
	decimal "github.com/shopspring/decimal"

	mock "github.com/stretchr/testify/mock"
)

// WalletRepo is an autogenerated mock type for the WalletRepo type
type WalletRepo struct {
	mock.Mock
}

// CreateOne provides a mock function with given fields: ctx, transaction
func (_m *WalletRepo) CreateOne(ctx context.Context, transaction *entities.WalletTransaction) (*entities.WalletTransaction, error) {
	ret := _m.Called(ctx, transaction)

	var r0 *entities.WalletTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.WalletTransaction) (*entities.WalletTransaction, error)); ok {
		return rf(ctx, transaction)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.WalletTransaction) *entities.WalletTransaction); ok {
		r0 = rf(ctx, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.WalletTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.WalletTransaction) error); ok {
		r1 = rf(ctx, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalAmount provides a mock function with given fields: ctx, userId
func (_m *WalletRepo) GetTotalAmount(ctx context.Context, userId int64) (decimal.Decimal, error) {
	ret := _m.Called(ctx, userId)

	var r0 decimal.Decimal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (decimal.Decimal, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) decimal.Decimal); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(decimal.Decimal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTransactions provides a mock function with given fields: ctx, userId, page, limit
func (_m *WalletRepo) ListTransactions(ctx context.Context, userId int64, page int, limit int) ([]*entities.WalletTransaction, error) {
	ret := _m.Called(ctx, userId, page, limit)

	var r0 []*entities.WalletTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, int) ([]*entities.WalletTransaction, error)); ok {
		return rf(ctx, userId, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, int) []*entities.WalletTransaction); ok {
		r0 = rf(ctx, userId, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.WalletTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int, int) error); ok {
		r1 = rf(ctx, userId, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWalletRepo creates a new instance of WalletRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWalletRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *WalletRepo {
	mock := &WalletRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}