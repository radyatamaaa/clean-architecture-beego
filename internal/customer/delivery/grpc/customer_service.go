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

func (v *CustomerService) mappingResultGetCustomers(r []domain.Customer) []*GetCustomersDto {
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
			CreatedAt:     r[i].CreatedAt.String(),
			UpdatedAt:     r[i].UpdatedAt.String(),
		}
	}

	return res
}

func (p CustomerService) mustEmbedUnimplementedCustomerServiceServer() {
	panic("implement me")
}
