package usecase

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/logger"
	"context"
	"database/sql"
	"time"
)

type productUseCase struct {
	contextTimeout    time.Duration
	productRepository domain.ProductRepository
	log logger.Logger
}



func NewProductUseCase(timeout time.Duration, ur domain.ProductRepository,log logger.Logger) domain.ProductUseCase {
	return &productUseCase{
		productRepository: ur,
		contextTimeout:    timeout,
		log : log,
	}
}

func (p productUseCase) GetProducts(c context.Context, limit, offset int) ([]domain.ProductObjectResponse, error) {
	log := "internal.product.usecase.productUseCase.GetProducts: %s"
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	var pList []domain.ProductObjectResponse
	result ,err := p.productRepository.Fetch(ctx, limit, offset)
	if err != nil{
		p.log.Error(log,err.Error())
		return nil, err
	}
	for _, v := range result {
		pList = append(pList, v.ToProductResponse())
	}

	return pList, nil

}

func (p productUseCase) GetProductById(c context.Context, id uint) (*domain.ProductObjectResponse, error) {
	log := "internal.product.usecase.productUseCase.GetProductById: %s"
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	result, err := p.productRepository.FindByID(ctx, id)
	if err != nil{
		p.log.Error(log,err.Error())
		return nil, err
	}

	product := result.ToProductResponse()
	return &product, err
}

func (p productUseCase) SaveProduct(c context.Context, body domain.ProductStoreRequest) error {
	log := "internal.product.usecase.productUseCase.SaveProduct: %s"
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	err := p.productRepository.Store(ctx, domain.Product{
		ProductName: body.Name,
		Price:       sql.NullFloat64{Float64: *body.Price, Valid: body.Price != nil},
		ActiveSale:  false,
		Stock:       sql.NullInt64{Int64: *body.Stock, Valid: body.Stock != nil},
	})

	if err != nil{
		p.log.Error(log,err.Error())
		return err
	}

	return nil
}

func (p productUseCase) UpdateProduct(c context.Context, body domain.ProductUpdateRequest) error {
	log := "internal.product.usecase.productUseCase.UpdateProduct: %s"
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	err := p.productRepository.Update(ctx, domain.Product{
		Id:          body.Id,
		ProductName: body.Name,
		Price:       sql.NullFloat64{Float64: *body.Price, Valid: body.Price != nil},
		ActiveSale:  false,
		Stock:       sql.NullInt64{Int64: *body.Stock, Valid: body.Stock != nil},
	})
	if err != nil{
		p.log.Error(log,err.Error())
		return err
	}

	return nil
}

func (p productUseCase) DeleteProduct(c context.Context, id int) error {
	log := "internal.product.usecase.productUseCase.DeleteProduct: %s"
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	err := p.productRepository.Delete(ctx, id)
	if err != nil{
		p.log.Error(log,err.Error())
		return err
	}
	return nil
}