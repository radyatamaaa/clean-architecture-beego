package http

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/helpers/converter_value"
	"clean-architecture-beego/pkg/helpers/response"
	"clean-architecture-beego/pkg/validator"
	"context"
	"errors"
	"net/http"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type CustomerHandler struct {
	beego.Controller
	response.ApiResponse
	CustomerUseCase domain.CustomerUseCase
}

func NewCustomerHandler(useCase domain.CustomerUseCase) {
	handler := &CustomerHandler{
		CustomerUseCase: useCase,
	}

	beego.Router("/api/v1/customer", handler, "get:GetCustomers")
	beego.Router("/api/v1/customer/:id", handler, "get:GetCustomerByID")
	beego.Router("/api/v1/customer", handler, "post:StoreCustomer")
	beego.Router("/api/v1/customer", handler, "put:UpdateCustomer")
	beego.Router("/api/v1/customer/:id", handler, "delete:DeleteCustomer")
}

func (h *CustomerHandler) GetCustomers() {

	var pageSize = 10
	var page = 0

	if parse, err := strconv.Atoi(h.Ctx.Input.Query("pageSize")); err == nil {
		pageSize = parse
	}
	if parse, err := strconv.Atoi(h.Ctx.Input.Query("page")); err == nil {
		page = parse
	}

	result, err := h.CustomerUseCase.GetCustomers(h.Ctx.Request.Context(), pageSize, page)

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

func (h *CustomerHandler) StoreCustomer() {

	var request domain.CustomerStoreRequest

	if err := h.BindJSON(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := validator.Validate.ValidateStruct(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := h.CustomerUseCase.SaveCustomer(h.Ctx.Request.Context(), request); err != nil {
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

func (h *CustomerHandler) UpdateCustomer() {
	var request domain.CustomerUpdateRequest

	if err := h.BindJSON(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := validator.Validate.ValidateStruct(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := h.CustomerUseCase.UpdateCustomer(h.Ctx.Request.Context(), request); err != nil {
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

func (h *CustomerHandler) DeleteCustomer() {

	id := converter_value.StringToInt(h.Ctx.Input.Param("id"))

	err := h.CustomerUseCase.DeleteCustomer(h.Ctx.Request.Context(), uint(id))

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.ErrorResponse(h.Ctx, http.StatusRequestTimeout, response.RequestTimeout, err)
			return
		}
		h.ErrorResponse(h.Ctx, http.StatusInternalServerError, response.ServerError, err)
		return
	}

	h.Ok(h.Ctx, id)
	return

}

func (h *CustomerHandler) GetCustomerByID() {

	id := converter_value.StringToInt(h.Ctx.Input.Param("id"))

	result, err := h.CustomerUseCase.GetCustomerById(h.Ctx.Request.Context(), uint(id))

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
