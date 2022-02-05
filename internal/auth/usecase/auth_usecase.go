package usecase

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthUseCase struct {
	contextTimeout time.Duration
}

func NewAuthUseCase(timeout time.Duration) domain.AuthUseCase {
	return &AuthUseCase{
		contextTimeout: timeout,
	}
}

func (p AuthUseCase) SignIn(c context.Context, username, password string) (string, error) {

	claims := &jwt.StandardClaims{
		ExpiresAt: 1800,
		Subject:   username,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret, err := beego.AppConfig.String("JWTSecret")
	if err != nil {
		panic(err)
	}

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
