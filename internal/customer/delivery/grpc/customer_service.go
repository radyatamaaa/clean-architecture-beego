package grpc

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"strconv"

	"google.golang.org/protobuf/runtime/protoimpl"
)

type CustomerService struct {
	CustomerUseCase domain.CustomerUseCase
}

func NewCustomerService(customerUseCase domain.CustomerUseCase) *CustomerService {
	return &CustomerService{
		CustomerUseCase: customerUseCase,
	}
}

func (p CustomerService) GetCustomers(ctx context.Context, params *GetCustomersParams) (*GetCustomersResult, error) {
	result := new(GetCustomersResult)
	if ctx == nil {
		ctx = context.Background()
	}

	// default
	var limit = 10
	var offset = 0

	limitParam := params.Limit
	offsetParam := params.Offset

	if parse, err := strconv.Atoi(limitParam); err == nil {
		limit = parse
	}
	if parse, err := strconv.Atoi(offsetParam); err == nil {
		offset = parse
	}
	res, err := p.CustomerUseCase.GetCustomers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	result.Data = p.mappingResultGetCustomers(res)
	result.Message = "success"

	return result, nil

}

func (p CustomerService) GetCustomerById(ctx context.Context, params *GetCustomerByIdParams) (*GetCustomerByIdResult, error) {
	result := new(GetCustomerByIdResult)
	if ctx == nil {
		ctx = context.Background()
	}

	id := uint(params.Id)

	res, err := p.CustomerUseCase.GetCustomerById(ctx, id)
	if err != nil {
		return nil, err
	}

	result.Data = p.mappingResultGetCustomerById(res)
	result.Message = "success"

	return result, nil

}

func (p CustomerService) StoreCustomer(ctx context.Context, req *CustomerStoreRequest) (*CustomerMessageResult, error) {
	result := new(CustomerMessageResult)
	if ctx == nil {
		ctx = context.Background()
	}

	var body domain.CustomerStoreRequest
	body.CustomerName = req.CustomerName
	body.Email = req.Email
	body.Phone = req.Phone
	body.Address = req.Address

	err := p.CustomerUseCase.SaveCustomer(ctx, body)
	if err != nil {
		return nil, err
	}

	result.Message = "success"

	return result, nil
}

func (p CustomerService) UpdateCustomer(ctx context.Context, req *CustomerUpdateRequest) (*CustomerMessageResult, error) {
	result := new(CustomerMessageResult)
	if ctx == nil {
		ctx = context.Background()
	}

	var body domain.CustomerUpdateRequest
	body.Id = uint(req.Id)
	body.CustomerName = req.CustomerName
	body.Email = req.Email
	body.Phone = req.Phone
	body.Address = req.Address

	err := p.CustomerUseCase.UpdateCustomer(ctx, body)
	if err != nil {
		return nil, err
	}

	result.Message = "success"

	return result, nil
}

func (p CustomerService) DeleteCustomer(ctx context.Context, params *GetCustomerByIdParams) (*CustomerMessageResult, error) {
	result := new(CustomerMessageResult)
	if ctx == nil {
		ctx = context.Background()
	}

	id := uint(params.Id)

	err := p.CustomerUseCase.DeleteCustomer(ctx, id)
	if err != nil {
		return nil, err
	}

	result.Message = "success"

	return result, nil
}

func (v *CustomerService) mappingResultGetCustomers(r []domain.CustomerObjectResponse) []*GetCustomersDto {
	res := make([]*GetCustomersDto, len(r))
	for i := range r {
		res[i] = &GetCustomersDto{
			state:         protoimpl.MessageState{},
			sizeCache:     0,
			unknownFields: nil,
			Id:            int32(r[i].Id),
			CustomerName:  r[i].CustomerName,
			Phone:         r[i].Phone,
			Email:         r[i].Email,
			Address:       r[i].Address,
		}
	}

	return res
}

func (v *CustomerService) mappingResultGetCustomerById(r *domain.CustomerObjectResponse) *GetCustomersDto {
	res := &GetCustomersDto{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Id:            int32(r.Id),
		CustomerName:  r.CustomerName,
		Phone:         r.Phone,
		Email:         r.Email,
		Address:       r.Address,
	}

	return res
}
func (p CustomerService) mustEmbedUnimplementedCustomerServiceServer() {
	panic("implement me")
}
