package repository

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/database"
	"context"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}


func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (p userRepository) Fetch(ctx context.Context,limit int, offset int) ([]domain.User, error) {
	var entities []domain.User
	paginator := database.NewPaginator(p.DB, offset, limit, &entities)
	return entities, paginator.Find(ctx).Error
}

func (p userRepository) FindByID(ctx context.Context,id uint) (domain.User, error) {
	var entity domain.User
	return entity, p.DB.WithContext(ctx).First(&entity, "id =?", id).Error
}

func (p userRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	var entity domain.User
	return entity, p.DB.WithContext(ctx).First(&entity, "username =?", username).Error
}

func (p userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var entity domain.User
	return entity, p.DB.WithContext(ctx).First(&entity, "email =?", email).Error
}

func (p userRepository) Update(ctx context.Context, user domain.User) error {
	return p.DB.WithContext(ctx).Updates(&user).Error
}

func (p userRepository) Store(ctx context.Context, user domain.User) error {
	return p.DB.WithContext(ctx).Create(&user).Error
}

func (p userRepository) Delete(ctx context.Context, id int) error {
	return p.DB.WithContext(ctx).Exec("delete from products where id =?", id).Error
}
