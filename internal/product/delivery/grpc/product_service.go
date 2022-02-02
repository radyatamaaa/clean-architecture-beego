package grpc

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"google.golang.org/protobuf/runtime/protoimpl"
	"strconv"
)

type ProductService struct {
	ProductUsecase domain.ProductUseCase
}

func NewProductService(productUsecase domain.ProductUseCase) *ProductService {
	return &ProductService{
		ProductUsecase:              productUsecase,
	}
}

func (p ProductService) GetProductsAPI(ctx context.Context, request *GetProductsRequest) (*GetProductsResponse, error) {
	panic("implement me")
}

func (p ProductService) GetProducts(ctx context.Context, param *GetProductsRequest) (*GetProductsResponse, error) {
	result := new(GetProductsResponse)
	if ctx == nil {
		ctx = context.Background()
	}

	// default
	var limit = 10
	var offset = 0

	if parse, err := strconv.Atoi(param.Limit); err == nil {
		limit = parse
	}
	if parse, err := strconv.Atoi(param.Offset); err == nil {
		offset = parse
	}
	res, err := p.ProductUsecase.GetProducts(ctx,limit, offset)
	if err != nil {
		return nil, err
	}

	result.Data = p.mappingResultGetProducts(res)
	result.Message = "success"
	return result,nil
}
func (v *ProductService)mappingResultGetProducts(r []domain.Product) []*GetProductsResponseDto {
	res := make([]*GetProductsResponseDto,len(r))
	for i := range r {
		res[i] = &GetProductsResponseDto{
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



