package grpc

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/helpers/converter_value"
	"clean-architecture-beego/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/protobuf/runtime/protoimpl"
	"strconv"
)

type ProductService struct {
	ProductUseCase domain.ProductUseCase
	log logger.Logger
}

func NewProductService(	productUseCase domain.ProductUseCase,log logger.Logger) *ProductService {
	return &ProductService{
		ProductUseCase:productUseCase,
		log: log,
	}
}

func (p ProductService) GetProducts(ctx context.Context, params *GetProductsParams) (*GetProductsResult, error) {
	log := "internal.delivery.grpc.ProductService.GetProducts: %s"

	result := new(GetProductsResult)
	if ctx == nil {
		ctx = context.Background()
	}

	v := ctx.Value("JWT_PAYLOAD")
	fmt.Println(v)

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
		p.log.Error(log,err.Error())
		return nil,err
	}

	result.Data = p.mappingResultGetProducts(res)
	result.Message = "success"

	return result,nil

}

func (v *ProductService)mappingResultGetProducts(r []domain.ProductObjectResponse) []*GetProductsDto {
	res := make([]*GetProductsDto,len(r))
	for i := range r {
		res[i] = &GetProductsDto{
			state:         protoimpl.MessageState{},
			sizeCache:     0,
			unknownFields: nil,
			Id:           int32(r[i].Id) ,
			ProductName:   r[i].ProductName,
			Price:       float32(converter_value.FloatNUllableToFloat(r[i].Price))  ,
			ActiveSale:    r[i].ActiveSale,
			Stock:       int32(converter_value.IntNullableToInt64(r[i].Stock)),
		}
	}

	return res
}

func (p ProductService) mustEmbedUnimplementedProductServiceServer() {
	panic("implement me")
}
