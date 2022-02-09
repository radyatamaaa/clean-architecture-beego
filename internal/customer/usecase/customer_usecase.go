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

func (p customerUseCase) GetCustomers(c context.Context, limit, offset int) ([]domain.CustomerObjectResponse, error) {
	var customerList []domain.CustomerObjectResponse
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	if result, err := p.customerRepository.Fetch(ctx, limit, offset); err != nil {
		return nil, err
	} else {
		for _, v := range result {
			customerList = append(customerList, v.ToCustomerResponse())
		}
	}
	return customerList, nil
}

func (p customerUseCase) GetCustomerById(c context.Context, id uint) (*domain.CustomerObjectResponse, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	result, err := p.customerRepository.FindByID(ctx, id)
	product := result.ToCustomerResponse()
	return &product, err
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

func (p customerUseCase) DeleteCustomer(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	return p.customerRepository.Delete(ctx, id)
}
