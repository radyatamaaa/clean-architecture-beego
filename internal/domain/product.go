package domain

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          uint            `gorm:"primarykey;autoIncrement:true"`
	ProductName string          `gorm:"type:varchar(50);column:product_name"`
	Price       sql.NullFloat64 `gorm:"column:product_price"`
	ActiveSale  bool            `gorm:"column:active_sale"`
	Stock       sql.NullInt64   `gorm:"column:stock"`
	CreatedAt   time.Time       `gorm:"column:created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at"`
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
	GetProducts(limit, offset int) ([]Product, error)
	GetProductById(id uint) (*Product, error)
	SaveProduct(body ProductStoreRequest) error
	UpdateProduct(body ProductUpdateRequest) error
}

type ProductRepository interface {
	Fetch(limit int, offset int) ([]Product, error)
	FindByID(id uint) (Product, error)
	Update(product Product) error
	Store(product Product) error
	Delete(id int) error
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
