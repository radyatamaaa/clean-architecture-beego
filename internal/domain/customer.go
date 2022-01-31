package domain

import (
	"time"
)

type Customer struct {
	Id           uint      `gorm:"primarykey;autoIncrement:true"`
	CustomerName string    `gorm:"type:varchar(100);column:customer_name"`
	Phone        string    `gorm:"type:varchar(14);column:phone;unique"`
	Email        string    `gorm:"type:varchar(50);column:email;unique"`
	Address      string    `gorm:"type:varchar(150);column:address"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
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
	GetCustomer(limit, offset int) ([]Order, error)
	GetCustomerById(body OrderStoreRequest) error
	SaveCustomer(body OrderStoreRequest) error
	UpdateCustomer(body OrderUpdateRequest) error
}

type CustomerRepository interface {
	Fetch(limit int, offset int) ([]Customer, error)
	FindByID(id string) (Customer, error)
	Update(order Customer) error
	Store(order Customer) error
	Delete(id int) error
}
