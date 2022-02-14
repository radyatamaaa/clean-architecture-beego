package repository

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/database"
	"clean-architecture-beego/pkg/logger"
	"context"
	"fmt"
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

func (p productRepository) Fetch(ctx context.Context,limit int, offset int) ([]domain.Product, error,string) {
	log := "internal.product.repository.productRepository.Fetch: %s"

	var entities []domain.Product
	paginator := database.NewPaginator(p.DB, offset, limit, &entities)
	err := paginator.Find(ctx).Error
	if err != nil{
		return nil,err, fmt.Sprintf(log+"\n", err.Error())
	}
	return entities,nil,""
}

func (p productRepository) FindByID(ctx context.Context,id uint) (domain.Product, error,string) {
	log := "internal.product.repository.productRepository.FindByID: %s"

	var entity domain.Product
	err := p.DB.WithContext(ctx).First(&entity, "id =?", id).Error
	if err != nil{
		return domain.Product{},err, fmt.Sprintf(log+"\n", err.Error())
	}
	return entity,nil,""
}

func (p productRepository) Update(ctx context.Context,product domain.Product) (error,string) {
	log := "internal.product.repository.productRepository.Update: %s"
	err := p.DB.WithContext(ctx).Updates(&product).Error
	if err != nil{
		return err , fmt.Sprintf(log+"\n", err.Error())
	}
	return nil,""
}

func (p productRepository) Store(ctx context.Context,product domain.Product) (error,string) {
	log := "internal.product.repository.productRepository.Store: %s"
	err := p.DB.WithContext(ctx).Create(&product).Error
	if err != nil{
		return err , fmt.Sprintf(log+"\n", err.Error())
	}
	return nil,""
}

func (p productRepository) Delete(ctx context.Context,id int) (error,string) {
	log := "internal.product.repository.productRepository.Delete: %s"
	err := p.DB.WithContext(ctx).Exec("delete from products where id =?", id).Error
	if err != nil{
		return err , fmt.Sprintf(log+"\n", err.Error())
	}
	return nil,""
}
