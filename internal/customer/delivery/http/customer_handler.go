package http

import (
	"clean-architecture-beego/internal/domain"
	"context"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type CustomerHandler struct {
	beego.Controller
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

	ctx := h.Ctx.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// default
	var limit = 10
	var offset = 0

	limitParam := h.Ctx.Input.Param("limit")
	offsetParam := h.Ctx.Input.Param("offset")

	if parse, err := strconv.Atoi(limitParam); err == nil {
		limit = parse
	}
	if parse, err := strconv.Atoi(offsetParam); err == nil {
		offset = parse
	}
	result, err := h.CustomerUseCase.GetCustomers(ctx, limit, offset)
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
		"data":    result,
	}
	if err := h.ServeJSON(); err != nil {
		return
	}
	return
}

func (h *CustomerHandler) StoreCustomer() {

}

func (h *CustomerHandler) UpdateCustomer() {

}

func (h *CustomerHandler) DeleteCustomer() {

}

func (h *CustomerHandler) GetCustomerByID() {

}
