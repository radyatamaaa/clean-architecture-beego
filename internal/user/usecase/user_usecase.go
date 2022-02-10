package usecase

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type userUseCase struct {
	contextTimeout time.Duration
	userRepository domain.UserRepository
}

func NewUserUseCase(timeout time.Duration, ur domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepository: ur,
		contextTimeout: timeout,
	}
}

func (u userUseCase) Login(c context.Context, username, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	result, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result, err := u.userRepository.FindByEmail(ctx, email)
			if err != nil {
				return nil, err
			}
			return &result, nil
		} else {
			return nil, err
		}
	}
	return &result, nil
}

func (u userUseCase) GetUsers(c context.Context, limit, offset int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.Fetch(ctx, limit, offset)
}

func (u userUseCase) GetUserById(ctx context.Context, id uint) (*domain.User, error) {
	panic("implement me")
}

func (u userUseCase) SaveUser(c context.Context, body domain.UserStoreRequest) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.userRepository.Store(ctx, domain.User{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})
}

func (u userUseCase) UpdateUser(ctx context.Context, body domain.UserUpdateRequest) error {
	panic("implement me")
}
