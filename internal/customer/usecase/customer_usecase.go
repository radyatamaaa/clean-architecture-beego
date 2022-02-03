package usecase

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"time"
)

type customerUseCase struct {
	contextTimeout     time.Duration
	customerRepository domain.CustomerRepository
}

func NewCustomerUseCase(timeout time.Duration, ur domain.CustomerRepository) domain.CustomerUseCase {
	return &customerUseCase{
		customerRepository: ur,
		contextTimeout:     timeout,
	}
}

func (p customerUseCase) GetCustomers(c context.Context, limit, offset int) ([]domain.Customer, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.customerRepository.Fetch(ctx, limit, offset)
}

func (p customerUseCase) GetCustomerById(c context.Context, id uint) (*domain.Customer, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	result, err := p.customerRepository.FindByID(ctx, id)
	return &result, err
}

func (p customerUseCase) SaveCustomer(c context.Context, body domain.CustomerStoreRequest) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	return p.customerRepository.Store(ctx, domain.Customer{
		CustomerName: body.CustomerName,
		Phone:        body.Phone,
		Email:        body.Email,
		Address:      body.Address,
	})
}

func (p customerUseCase) UpdateCustomer(c context.Context, body domain.CustomerUpdateRequest) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.customerRepository.Update(ctx, domain.Customer{
		Id:           body.Id,
		CustomerName: body.CustomerName,
		Phone:        body.Phone,
		Email:        body.Email,
		Address:      body.Address,
	})
}
