package domain

import (
	"context"
	"database/sql"
	"time"
)
type ProductTest struct {
	Id          uint            `gorm:"primarykey;autoIncrement:true"`
	ProductName string          `gorm:"type:varchar(50);column:product_name"`
	Price       float64 `gorm:"column:product_price"`
	ActiveSale  bool            `gorm:"column:active_sale"`
	Stock       int   `gorm:"column:stock"`
	CreatedAt   time.Time       `gorm:"column:created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at"`
}
type Product struct {
	Id          uint            `json:"id" gorm:"primarykey;autoIncrement:true"`
	ProductName string          `json:"product_name" gorm:"type:varchar(50);column:product_name"`
	Price       sql.NullFloat64 `json:"product_price" gorm:"column:product_price"`
	ActiveSale  bool            `json:"active_sale" gorm:"column:active_sale"`
	Stock       sql.NullInt64   `json:"stock" gorm:"column:stock"`
	CreatedAt   time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"column:updated_at"`
}

func (p *Product) TableName() string {
	return "products"
}

type ProductObjectResponse struct {
	Id          uint     `json:"id"`
	ProductName string   `json:"product_name"`
	Price       *float64 `json:"price"`
	ActiveSale  bool     `json:"active_sale"`
	Stock       *int64   `json:"stock"`
}

type ProductStoreRequest struct {
	Name  string   `json:"name"`
	Price *float64 `json:"price"`
	Stock *int64   `json:"stock"`
}

type ProductUpdateRequest struct {
	Id    uint     `json:"id"`
	Name  string   `json:"name"`
	Price *float64 `json:"price"`
	Stock *int64   `json:"stock"`
}

type ProductUseCase interface {
	GetProducts(ctx context.Context,limit, offset int) ([]Product, error)
	GetProductById(ctx context.Context,id uint) (*Product, error)
	SaveProduct(ctx context.Context,body ProductStoreRequest) error
	UpdateProduct(ctx context.Context,body ProductUpdateRequest) error
}

type ProductRepository interface {
	Fetch(ctx context.Context,limit int, offset int) ([]Product, error)
	FindByID(ctx context.Context,id uint) (Product, error)
	Update(ctx context.Context,product Product) error
	Store(ctx context.Context,product Product) error
	Delete(ctx context.Context,id int) error
}

func (p Product) ToProductResponse() ProductObjectResponse {
	var price *float64
	var stock *int64

	if p.Price.Valid {
		price = &p.Price.Float64
	}
	if p.Stock.Valid {
		stock = &p.Stock.Int64
	}

	return ProductObjectResponse{
		Id:          p.Id,
		ProductName: p.ProductName,
		Price:       price,
		ActiveSale:  p.ActiveSale,
		Stock:       stock,
	}
}

func (p ProductTest) ToProduct() Product {
	return Product{
		Id:          p.Id,
		ProductName: p.ProductName,
		Price:       sql.NullFloat64{Float64: p.Price},
		ActiveSale:  false,
		Stock:       sql.NullInt64{
			Int64: int64(p.Stock),
			Valid: false,
		},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}