package repository

import (
	"clean-architecture-beego/database"
	"clean-architecture-beego/internal/domain"
	"context"

	"gorm.io/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{
		DB: db,
	}
}

func (p customerRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Customer, error) {
	var entities []domain.Customer
	paginator := database.NewPaginator(p.DB, offset, limit, &entities)
	err := paginator.Find(ctx).Error
	return entities, err
}

func (p customerRepository) FindByID(ctx context.Context, id uint) (domain.Customer, error) {
	var entity domain.Customer
	return entity, p.DB.WithContext(ctx).First(&entity, "id =?", id).Error
}

func (p customerRepository) Update(ctx context.Context, customer domain.Customer) error {
	return p.DB.WithContext(ctx).Updates(&customer).Error
}

func (p customerRepository) Store(ctx context.Context, customer domain.Customer) error {
	return p.DB.WithContext(ctx).Create(&customer).Error
}

func (p customerRepository) Delete(ctx context.Context, id uint) error {
	return p.DB.WithContext(ctx).Exec("delete from customers where id =?", id).Error
}
