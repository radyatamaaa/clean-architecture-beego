package usecase

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"database/sql"
	"time"
)

type productUseCase struct {
	contextTimeout    time.Duration
	productRepository domain.ProductRepository
}



func NewProductUseCase(timeout time.Duration, ur domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{
		productRepository: ur,
		contextTimeout:    timeout,
	}
}

func (p productUseCase) GetProducts(c context.Context, limit, offset int) ([]domain.ProductObjectResponse, error) {
	var pList []domain.ProductObjectResponse
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	if result, err := p.productRepository.Fetch(ctx, limit, offset); err != nil {
		return nil, err
	} else {
		for _, v := range result {
			pList = append(pList, v.ToProductResponse())
		}
	}
	return pList, nil
}

func (p productUseCase) GetProductById(c context.Context, id uint) (*domain.ProductObjectResponse, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	result, err := p.productRepository.FindByID(ctx, id)
	product := result.ToProductResponse()
	return &product, err
}

func (p productUseCase) SaveProduct(c context.Context, body domain.ProductStoreRequest) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	return p.productRepository.Store(ctx, domain.Product{
		ProductName: body.Name,
		Price:       sql.NullFloat64{Float64: *body.Price, Valid: body.Price != nil},
		ActiveSale:  false,
		Stock:       sql.NullInt64{Int64: *body.Stock, Valid: body.Stock != nil},
	})
}

func (p productUseCase) UpdateProduct(c context.Context, body domain.ProductUpdateRequest) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.productRepository.Update(ctx, domain.Product{
		Id:          body.Id,
		ProductName: body.Name,
		Price:       sql.NullFloat64{Float64: *body.Price, Valid: body.Price != nil},
		ActiveSale:  false,
		Stock:       sql.NullInt64{Int64: *body.Stock, Valid: body.Stock != nil},
	})
}

func (p productUseCase) DeleteProduct(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	return p.productRepository.Delete(ctx, id)
}