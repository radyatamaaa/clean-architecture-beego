package grpc

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"google.golang.org/protobuf/runtime/protoimpl"
	"strconv"
)

type ProductService struct {
	ProductUseCase domain.ProductUseCase
}

func NewProductService(	productUseCase domain.ProductUseCase) *ProductService {
	return &ProductService{
		ProductUseCase:productUseCase,
	}
}

func (p ProductService) GetProducts(ctx context.Context, params *GetProductsParams) (*GetProductsResult, error) {
	result := new(GetProductsResult)
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
	res, err := p.ProductUseCase.GetProducts(ctx,limit, offset)
	if err != nil{
		return nil, err
	}

	result.Data = p.mappingResultGetProducts(res)
	result.Message = "success"

	return result,nil

}

func (v *ProductService)mappingResultGetProducts(r []domain.Product) []*GetProductsDto {
	res := make([]*GetProductsDto,len(r))
	for i := range r {
		res[i] = &GetProductsDto{
			state:         protoimpl.MessageState{},
			sizeCache:     0,
			unknownFields: nil,
			Id:           int32(r[i].Id) ,
			ProductName:   r[i].ProductName,
			Price:       float32(r[i].Price.Float64)  ,
			ActiveSale:    r[i].ActiveSale,
			Stock:       int32(r[i].Stock.Int64),
			CreatedAt:     r[i].CreatedAt.String(),
			UpdatedAt:     r[i].UpdatedAt.String(),
		}
	}

	return res
}

func (p ProductService) mustEmbedUnimplementedProductServiceServer() {
	panic("implement me")
}
