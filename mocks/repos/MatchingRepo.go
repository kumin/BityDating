// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/kumin/AndPadDating/entities"
	mock "github.com/stretchr/testify/mock"
)

// MatchingRepo is an autogenerated mock type for the MatchingRepo type
type MatchingRepo struct {
	mock.Mock
}

// CreateOne provides a mock function with given fields: ctx, matching
func (_m *MatchingRepo) CreateOne(ctx context.Context, matching *entities.UserMatching) (*entities.UserMatching, error) {
	ret := _m.Called(ctx, matching)

	var r0 *entities.UserMatching
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserMatching) (*entities.UserMatching, error)); ok {
		return rf(ctx, matching)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserMatching) *entities.UserMatching); ok {
		r0 = rf(ctx, matching)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserMatching)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.UserMatching) error); ok {
		r1 = rf(ctx, matching)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOne provides a mock function with given fields: ctx, userId, partnerId
func (_m *MatchingRepo) DeleteOne(ctx context.Context, userId int64, partnerId int64) error {
	ret := _m.Called(ctx, userId, partnerId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userId, partnerId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListMatching provides a mock function with given fields: ctx, userId, page, limit
func (_m *MatchingRepo) ListMatching(ctx context.Context, userId int64, page int, limit int) ([]*entities.User, error) {
	ret := _m.Called(ctx, userId, page, limit)

	var r0 []*entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, int) ([]*entities.User, error)); ok {
		return rf(ctx, userId, page, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, int) []*entities.User); ok {
		r0 = rf(ctx, userId, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int, int) error); ok {
		r1 = rf(ctx, userId, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WhoILike provides a mock function with given fields: ctx, userId
func (_m *MatchingRepo) WhoILike(ctx context.Context, userId int64) ([]*entities.User, error) {
	ret := _m.Called(ctx, userId)

	var r0 []*entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*entities.User, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*entities.User); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WhoLikeMe provides a mock function with given fields: ctx, partnerId
func (_m *MatchingRepo) WhoLikeMe(ctx context.Context, partnerId int64) ([]*entities.User, error) {
	ret := _m.Called(ctx, partnerId)

	var r0 []*entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*entities.User, error)); ok {
		return rf(ctx, partnerId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*entities.User); ok {
		r0 = rf(ctx, partnerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, partnerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMatchingRepo creates a new instance of MatchingRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMatchingRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MatchingRepo {
	mock := &MatchingRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}