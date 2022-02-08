package mocks

import (
	"clean-architecture-beego/internal/domain"
	"context"

	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Customer, error) {
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

func (_m *Repository) FindByID(ctx context.Context, id uint) (domain.Customer, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Customer
	if rf, ok := ret.Get(0).(func(context.Context, uint) domain.Customer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Customer)
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

func (_m *Repository) Update(ctx context.Context, customer domain.Customer) error {
	ret := _m.Called(ctx, customer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Customer) error); ok {
		r0 = rf(ctx, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Repository) Store(ctx context.Context, customer domain.Customer) error {
	ret := _m.Called(ctx, customer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Customer) error); ok {
		r0 = rf(ctx, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Repository) Delete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, int(id))
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
