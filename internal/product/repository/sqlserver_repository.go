package repository

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/database"
	"clean-architecture-beego/pkg/logger"
	"context"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
	log logger.Logger
}

func NewProductRepository(db *gorm.DB,log logger.Logger) domain.ProductRepository {
	return &productRepository{
		DB: db,
		log: log,
	}
}

func (p productRepository) Fetch(ctx context.Context,limit int, offset int) ([]domain.Product, error) {
	log := "internal.product.repository.productRepository.Fetch: %s"

	var entities []domain.Product
	paginator := database.NewPaginator(p.DB, offset, limit, &entities)
	err := paginator.Find(ctx).Error
	if err != nil{
		p.log.Error(log,err.Error())
		return nil,err
	}
	return entities,nil
}

func (p productRepository) FindByID(ctx context.Context,id uint) (domain.Product, error) {
	log := "internal.product.repository.productRepository.FindByID: %s"

	var entity domain.Product
	err := p.DB.WithContext(ctx).First(&entity, "id =?", id).Error
	if err != nil{
		p.log.Error(log,err.Error())
		return domain.Product{},err
	}
	return entity,nil
}

func (p productRepository) Update(ctx context.Context,product domain.Product) error {
	log := "internal.product.repository.productRepository.Update: %s"
	err := p.DB.WithContext(ctx).Updates(&product).Error
	if err != nil{
		p.log.Error(log,err.Error())
		return err
	}
	return nil
}

func (p productRepository) Store(ctx context.Context,product domain.Product) error {
	log := "internal.product.repository.productRepository.Store: %s"
	err := p.DB.WithContext(ctx).Create(&product).Error
	if err != nil{
		p.log.Error(log,err.Error())
		return err
	}
	return nil
}

func (p productRepository) Delete(ctx context.Context,id int) error {
	log := "internal.product.repository.productRepository.Delete: %s"
	err := p.DB.WithContext(ctx).Exec("delete from products where id =?", id).Error
	if err != nil{
		p.log.Error(log,err.Error())
		return err
	}
	return nil
}
