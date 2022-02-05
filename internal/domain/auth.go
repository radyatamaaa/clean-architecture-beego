package domain

import "context"

type AuthUseCase interface {
	SignIn(ctx context.Context, username, password string) (string, error)
	// Verify(ctx context.Context, accessToken string) error
}

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
