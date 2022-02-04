package usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthUseCase struct {
	expireDuration time.Duration
}

type AuthClaims struct {
	*jwt.StandardClaims
}
