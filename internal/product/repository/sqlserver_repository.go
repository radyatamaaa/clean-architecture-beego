package repository

import (
	"clean-architecture-beego/database"
	"clean-architecture-beego/internal/domain"
	"context"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (p productRepository) Fetch(ctx context.Context,limit int, offset int) ([]domain.Product, error) {
	var entities []domain.Product
	paginator := database.NewPaginator(p.DB, offset, limit, &entities)
	err := paginator.Find(ctx).Error
	return entities, err
}

func (p productRepository) FindByID(ctx context.Context,id uint) (domain.Product, error) {
	var entity domain.Product
	return entity, p.DB.WithContext(ctx).First(&entity, "id =?", id).Error
}

func (p productRepository) Update(ctx context.Context,product domain.Product) error {
	return p.DB.WithContext(ctx).Updates(&product).Error
}

func (p productRepository) Store(ctx context.Context,product domain.Product) error {
	return p.DB.WithContext(ctx).Create(&product).Error
}

func (p productRepository) Delete(ctx context.Context,id int) error {
	return p.DB.WithContext(ctx).Exec("delete from products where id =?", id).Error
}
