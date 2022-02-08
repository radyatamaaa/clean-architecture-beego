package mocks

import (
	"clean-architecture-beego/internal/domain"
	"context"

	mock "github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (_m *Usecase) GetCustomers(ctx context.Context, limit, offset int) ([]domain.Customer, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []domain.Customer
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []domain.Customer); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Usecase) GetCustomerById(ctx context.Context, id uint) (*domain.Customer, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Customer
	if rf, ok := ret.Get(0).(func(context.Context, uint) *domain.Customer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Usecase) SaveCustomer(ctx context.Context, body domain.CustomerStoreRequest) error {
	ret := _m.Called(ctx, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CustomerStoreRequest) error); ok {
		r0 = rf(ctx, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Usecase) UpdateCustomer(ctx context.Context, body domain.CustomerUpdateRequest) error {
	ret := _m.Called(ctx, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CustomerUpdateRequest) error); ok {
		r0 = rf(ctx, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
