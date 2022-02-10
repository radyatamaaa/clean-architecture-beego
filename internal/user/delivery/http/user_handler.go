package http

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/helpers/response"
	"clean-architecture-beego/pkg/jwt"
	"clean-architecture-beego/pkg/validator"
	"context"
	"errors"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	beego.Controller
	response.ApiResponse
	UserUseCase domain.UserUseCase
	JwtAuth     jwt.JWT
}

func NewUserHandler(useCase domain.UserUseCase, jwt jwt.JWT) {
	pHandler := &UserHandler{
		UserUseCase: useCase,
		JwtAuth:     jwt,
	}
	beego.Router("/token", pHandler, "post:RequestToken")
	beego.Router("/register", pHandler, "post:Register")
	beego.Router("/api/v1/refresh-token", pHandler, "post:RefreshToken")
}

func (h *UserHandler) URLMapping() {
	h.Mapping("GetUsers", h.GetUsers)
	h.Mapping("Register", h.Register)
	h.Mapping("RequestToken", h.RequestToken)
	h.Mapping("RefreshToken", h.RefreshToken)
}

func (h *UserHandler) GetUsers() {
	// default
	var pageSize = 10
	var page = 0

	if parse, err := strconv.Atoi(h.Ctx.Input.Query("pageSize")); err == nil {
		pageSize = parse
	}
	if parse, err := strconv.Atoi(h.Ctx.Input.Query("page")); err == nil {
		page = parse
	}

	result, err := h.UserUseCase.GetUsers(h.Ctx.Request.Context(), pageSize, page)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.ErrorResponse(h.Ctx, http.StatusRequestTimeout, response.RequestTimeout, err)
			return
		}
		h.ErrorResponse(h.Ctx, http.StatusInternalServerError, response.ServerError, err)
		return
	}

	h.Ok(h.Ctx, result)
	return
}

func (h *UserHandler) Register() {
	var request domain.UserStoreRequest

	if err := h.BindJSON(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := validator.Validate.ValidateStruct(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := h.UserUseCase.SaveUser(h.Ctx.Request.Context(), request); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.ErrorResponse(h.Ctx, http.StatusRequestTimeout, response.RequestTimeout, err)
			return
		}
		h.ErrorResponse(h.Ctx, http.StatusInternalServerError, response.ServerError, err)
		return
	}
	h.Ok(h.Ctx, request)
	return
}

func (h *UserHandler) RequestToken() {
	var request domain.UserLoginRequest

	if err := h.BindJSON(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := validator.Validate.ValidateStruct(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}

	if result, err := h.UserUseCase.Login(h.Ctx.Request.Context(), request.Username, request.Email); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.ErrorResponse(h.Ctx, http.StatusRequestTimeout, response.RequestTimeout, err)
			return
		}
		h.ErrorResponse(h.Ctx, http.StatusInternalServerError, response.ServerError, err)
		return
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password)); err != nil {
			h.ErrorResponse(h.Ctx, http.StatusBadRequest, response.InvalidCredentials, err)
			return
		}
		token, err := h.JwtAuth.Ctx(h.Ctx.Request.Context()).GenerateToken(jwt.Payload{"uid": result.Id, "username": result.Username})
		if err != nil {
			h.ErrorResponse(h.Ctx, http.StatusInternalServerError, response.ServerError, err)
			return
		}
		h.Ok(h.Ctx, token)
		return
	}
}

func (h *UserHandler) RefreshToken() {
	t, err := h.JwtAuth.Ctx(h.Ctx.Request.Context()).RefreshToken(h.Ctx.Request)

	logs.Error(err)
	if err != nil {
		h.ErrorResponse(h.Ctx, http.StatusInternalServerError, response.ServerError, err)
		return
	}
	h.Ok(h.Ctx, t)
	return
}
