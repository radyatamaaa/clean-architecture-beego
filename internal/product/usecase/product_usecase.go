package usecase

import (
	"clean-architecture-beego/internal/domain"
	"database/sql"
)

type productUseCase struct {
	productRepository domain.ProductRepository
}

func NewProductUseCase(ur domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{productRepository: ur}
}

func (p productUseCase) GetProducts(limit, offset int) ([]domain.Product, error) {
	return p.productRepository.Fetch(limit, offset)
}

func (p productUseCase) GetProductById(id uint) (*domain.Product, error) {
	result, err := p.productRepository.FindByID(id)
	return &result, err
}

func (p productUseCase) SaveProduct(body domain.ProductStoreRequest) error {
	return p.productRepository.Store(domain.Product{
		ProductName: body.Name,
		Price:       sql.NullFloat64{Float64: *body.Price, Valid: body.Price != nil},
		ActiveSale:  false,
		Stock:       sql.NullInt64{Int64: *body.Stock, Valid: body.Stock != nil},
	})
}

func (p productUseCase) UpdateProduct(body domain.ProductUpdateRequest) error {
	return p.productRepository.Update(domain.Product{
		Id:          body.Id,
		ProductName: body.Name,
		Price:       sql.NullFloat64{Float64: *body.Price, Valid: body.Price != nil},
		ActiveSale:  false,
		Stock:       sql.NullInt64{Int64: *body.Stock, Valid: body.Stock != nil},
	})
}
