// Code generated by mockery v1.0.0
package mocks

import (
	"clean-architecture-beego/internal/domain"
	"context"
	mock "github.com/stretchr/testify/mock"
)

// repository is an autogenerated mock type for the repository type
type Repository struct {
	mock.Mock
}

func (_m *Repository) Fetch(ctx context.Context,limit int, offset int) ([]domain.Product, error,string) {
	ret := _m.Called(ctx, limit,offset)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, int,int) []domain.Product); ok {
		r0 = rf(ctx, limit,offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int,int) error); ok {
		r1 = rf(ctx, limit,offset)
	} else {
		r1 = ret.Error(1)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(context.Context, int,int) string); ok {
		r2 = rf(ctx, limit,offset)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(string)
		}
	}

	return r0, r1,r2
}

func (_m *Repository) FindByID(ctx context.Context,id uint) (domain.Product, error,string) {
	ret := _m.Called(ctx, id)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, uint) domain.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(context.Context, uint) string); ok {
		r2 = rf(ctx, id)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(string)
		}
	}

	return r0, r1,r2
}

func (_m *Repository) Update(ctx context.Context,product domain.Product) (error,string) {
	ret := _m.Called(ctx,product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) error); ok {
		r0 = rf(ctx,product)
	} else {
		r0 = ret.Error(0)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, domain.Product) string); ok {
		r1 = rf(ctx, product)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(string)
		}
	}

	return r0,r1
}

func (_m *Repository) Store(ctx context.Context,product domain.Product) (error,string) {
	ret := _m.Called(ctx,product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) error); ok {
		r0 = rf(ctx,product)
	} else {
		r0 = ret.Error(0)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, domain.Product) string); ok {
		r1 = rf(ctx, product)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(string)
		}
	}

	return r0,r1
}

func (_m *Repository) Delete(ctx context.Context,id int) (error,string) {
	ret := _m.Called(ctx,id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx,id)
	} else {
		r0 = ret.Error(0)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, int) string); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(string)
		}
	}

	return r0,r1
}
