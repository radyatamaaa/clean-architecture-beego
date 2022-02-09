package domain

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint      `gorm:"primarykey;autoIncrement:true"`
	Username  string    `gorm:"type:varchar(50);column:username;unique"`
	Email     string    `gorm:"type:varchar(100);column:email;unique"`
	Password  string    `gorm:"type:varchar(200);column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (c *User) TableName() string {
	return "users"
}

func (c *User) BeforeCreate(tx *gorm.DB) (err error) {
	if bytes, err := bcrypt.GenerateFromPassword([]byte(c.Password), 10); err != nil {
		return err
	} else {
		c.Password = string(bytes)
	}
	return nil
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRefreshTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

type UserStoreRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUseCase interface {
	GetUsers(ctx context.Context, limit, offset int) ([]User, error)
	GetUserById(ctx context.Context, id uint) (*User, error)
	SaveUser(ctx context.Context, body UserStoreRequest) error
	UpdateUser(ctx context.Context, body UserUpdateRequest) error
	Login(ctx context.Context, username, email string) (*User, error)
}

type UserRepository interface {
	Fetch(ctx context.Context, limit int, offset int) ([]User, error)
	FindByID(ctx context.Context, id uint) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	Update(ctx context.Context, user User) error
	Store(ctx context.Context, user User) error
	Delete(ctx context.Context, id int) error
}
