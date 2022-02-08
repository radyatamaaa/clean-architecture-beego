package domain

import (
	"context"
	"time"
)

type CustomerTest struct {
	Id           uint      `gorm:"primarykey;autoIncrement:true"`
	CustomerName string    `gorm:"type:varchar(100);column:customer_name"`
	Phone        string    `gorm:"type:varchar(14);column:phone;unique"`
	Email        string    `gorm:"type:varchar(50);column:email;unique"`
	Address      string    `gorm:"type:varchar(150);column:address"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type Customer struct {
	Id           uint      `json:"id" gorm:"primarykey;autoIncrement:true"`
	CustomerName string    `json:"customer_name" gorm:"type:varchar(100);column:customer_name"`
	Phone        string    `json:"phone" gorm:"type:varchar(14);column:phone;unique"`
	Email        string    `json:"email" gorm:"type:varchar(50);column:email;unique"`
	Address      string    `json:"address" gorm:"type:varchar(150);column:address"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (c *Customer) TableName() string {
	return "customers"
}

type CustomerStoreRequest struct {
	CustomerName string `json:"customer_name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Address      string `json:"address"`
}

type CustomerUpdateRequest struct {
	Id           uint   `json:"id"`
	CustomerName string `json:"customer_name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Address      string `json:"address"`
}

type CustomerUseCase interface {
	GetCustomers(ctx context.Context, limit, offset int) ([]Customer, error)
	GetCustomerById(ctx context.Context, id uint) (*Customer, error)
	SaveCustomer(ctx context.Context, body CustomerStoreRequest) error
	UpdateCustomer(ctx context.Context, body CustomerUpdateRequest) error
	DeleteCustomer(ctx context.Context, id uint) error
}

type CustomerRepository interface {
	Fetch(ctx context.Context, limit int, offset int) ([]Customer, error)
	FindByID(ctx context.Context, id uint) (Customer, error)
	Update(ctx context.Context, product Customer) error
	Store(ctx context.Context, product Customer) error
	Delete(ctx context.Context, id uint) error
}

func (p CustomerTest) ToCustomer() Customer {
	return Customer{
		Id:           p.Id,
		CustomerName: p.CustomerName,
		Phone:        p.Phone,
		Email:        p.Email,
		Address:      p.Address,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
