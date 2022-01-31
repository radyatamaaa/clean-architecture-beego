package repository

import (
	"clean-architecture-beego/database"
	"clean-architecture-beego/internal/domain"
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

func (p productRepository) Fetch(limit int, offset int) ([]domain.Product, error) {
	var entities []domain.Product
	paginator := database.NewPaginator(p.DB, offset, limit, &entities)
	err := paginator.Find().Error
	return entities, err
}

func (p productRepository) FindByID(id uint) (domain.Product, error) {
	var entity domain.Product
	return entity, p.DB.First(&entity, "id =?", id).Error
}

func (p productRepository) Update(product domain.Product) error {
	return p.DB.Updates(&product).Error
}

func (p productRepository) Store(product domain.Product) error {
	return p.DB.Create(&product).Error
}

func (p productRepository) Delete(id int) error {
	return p.DB.Exec("delete from products where id =?", id).Error
}
