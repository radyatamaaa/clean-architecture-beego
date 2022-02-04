package http

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthHandler struct {
	beego.Controller
	AuthUseCase domain.AuthUseCase
}

func NewAuthHandler(useCase domain.AuthUseCase) {
	handler := &AuthHandler{
		AuthUseCase: useCase,
	}

	beego.Router("/api/v1/login", handler, "post:Login")
}

func (h *AuthHandler) Login() {
	ctx := h.Ctx.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var body domain.AuthLoginRequest

	json.Unmarshal(h.Ctx.Input.RequestBody, &body)
	accessToken, err := h.AuthUseCase.SignIn(ctx, body.Username, body.Password)
	if err != nil {
		h.Data["json"] = beego.M{
			"message": "internal server error",
			"error":   err,
		}
		if err := h.ServeJSON(); err != nil {
			return
		}
		return
	}
	h.Data["json"] = beego.M{
		"message": "success",
		"error":   nil,
		"data":    accessToken,
	}
	if err := h.ServeJSON(); err != nil {
		return
	}
	return

}
