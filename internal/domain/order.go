package domain

import (
	"database/sql"
	"time"
)

type Order struct {
	Id         uint            `gorm:"primarykey;autoIncrement:true"`
	CustomerId uint            `gorm:"column:customer_id"`
	OrderDate  sql.NullTime    `gorm:"column:order_date"`
	Qty        sql.NullInt64   `gorm:"column:qty"`
	Amount     sql.NullFloat64 `gorm:"column:amount"`
	CreatedAt  time.Time       `gorm:"column:created_at"`
	UpdatedAt  time.Time       `gorm:"column:updated_at"`
}

func (r *Order) TableName() string {
	return "orders"
}

type OrderObjectResponse struct {
	Id         uint     `json:"id"`
	CustomerId uint     `json:"customer_id"`
	OrderDate  string   `json:"order_date"`
	Qty        *int64   `json:"qty"`
	Amount     *float64 `json:"amount"`
}

type OrderStoreRequest struct {
	CustomerId uint    `json:"customer_id"`
	OrderDate  string  `json:"order_date"`
	Qty        int64   `json:"qty"`
	Amount     float64 `json:"amount"`
}

type OrderUpdateRequest struct {
	Id         uint     `json:"id"`
	CustomerId uint    `json:"customer_id"`
	OrderDate  string  `json:"order_date"`
	Qty        int64   `json:"qty"`
	Amount     float64 `json:"amount"`
}

type OrderUseCase interface {
	GetOrder(limit, offset int) ([]Order, error)
	GetOrderById(body OrderStoreRequest) error
	SaveOrder(body OrderStoreRequest) error
	UpdateOrder(body OrderUpdateRequest) error
}

type OrderRepository interface {
	Fetch(limit int, offset int) ([]Order, error)
	FindByID(id string) (Order, error)
	Update(order Order) error
	Store(order Order) error
	Delete(id int) error
}

func (r Order) ToOrderResponse() OrderObjectResponse {
	var orderDate string
	var qty *int64
	var amount *float64

	if r.OrderDate.Valid {
		orderDate = r.OrderDate.Time.Format(time.RFC3339)
	}
	if r.Qty.Valid {
		qty = &r.Qty.Int64
	}
	if r.Amount.Valid {
		amount = &r.Amount.Float64
	}
	return OrderObjectResponse{
		Id:         r.Id,
		CustomerId: r.CustomerId,
		OrderDate:  orderDate,
		Qty:        qty,
		Amount:     amount,
	}
}
