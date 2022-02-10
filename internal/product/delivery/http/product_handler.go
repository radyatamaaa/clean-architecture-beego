package http

import (
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/helpers/converter_value"
	"clean-architecture-beego/pkg/helpers/response"
	"clean-architecture-beego/pkg/logger"
	"clean-architecture-beego/pkg/validator"
	"context"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	beego.Controller
	response.ApiResponse
	ProductUseCase domain.ProductUseCase
	log logger.Logger
}

func NewProductHandler(useCase domain.ProductUseCase,	log logger.Logger) {
	pHandler := &ProductHandler{
		log:log,
		ProductUseCase: useCase,
	}
	beego.Router("/api/v1/products", pHandler, "get:GetProducts")
	beego.Router("/api/v1/product/:id", pHandler, "get:GetProductByID")
	beego.Router("/api/v1/product", pHandler, "post:StoreProduct")
	beego.Router("/api/v1/product", pHandler, "put:UpdateProduct")
	beego.Router("/api/v1/product/:id", pHandler, "delete:DeleteProduct")
}

//func (h *ProductHandler) URLMapping() {
//	h.Mapping("GetProducts", h.GetProducts)
//	h.Mapping("StoreProduct", h.StoreProduct)
//	h.Mapping("UpdateProduct", h.UpdateProduct)
//	h.Mapping("DeleteProduct", h.Delete)
//	h.Mapping("GetProductByID", h.GetProductByID)
//}

// GetProducts godoc
// @Summary Get all products
// @Tags Product
// @Produce json
// @Param pageSize query string false "page size"
// @Param page query string false "page"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 422 {object} response.ApiResponse{errors=[]response.Errors}
// @Failure 500 {object} response.ApiResponse
// @Router /v1/products [get]
func (h *ProductHandler) GetProducts() {
	log := "internal.delivery.http.ProductHandler.GetProducts: %s"
	// default
	var pageSize = 10
	var page = 0

	if parse, err := strconv.Atoi(h.Ctx.Input.Query("pageSize")); err == nil {
		pageSize = parse
	}
	if parse, err := strconv.Atoi(h.Ctx.Input.Query("page")); err == nil {
		page = parse
	}

	result, err := h.ProductUseCase.GetProducts(h.Ctx.Request.Context(), pageSize, page)

	if err != nil {
		h.log.Error(log,err.Error())
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

func (h *ProductHandler) StoreProduct() {
	log := "internal.delivery.http.ProductHandler.GetProducts: %s"
	var request domain.ProductStoreRequest

	if err := h.BindJSON(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := validator.Validate.ValidateStruct(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := h.ProductUseCase.SaveProduct(h.Ctx.Request.Context(), request); err != nil {
		h.log.Error(log,err.Error())
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

func (h *ProductHandler) UpdateProduct() {
	log := "internal.delivery.http.ProductHandler.GetProducts: %s"
	var request domain.ProductUpdateRequest

	if err := h.BindJSON(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := validator.Validate.ValidateStruct(&request); err != nil {
		h.ErrorResponse(h.Ctx, http.StatusUnprocessableEntity, response.ApiValidationError, err)
		return
	}
	if err := h.ProductUseCase.UpdateProduct(h.Ctx.Request.Context(), request); err != nil {
		h.log.Error(log,err.Error())
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

func (h *ProductHandler) DeleteProduct() {
	log := "internal.delivery.http.ProductHandler.GetProducts: %s"
	id := converter_value.StringToInt(h.Ctx.Input.Param("id"))

	err := h.ProductUseCase.DeleteProduct(h.Ctx.Request.Context(), id)

	if err != nil {
		h.log.Error(log,err.Error())
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

func (h *ProductHandler) GetProductByID() {
	log := "internal.delivery.http.ProductHandler.GetProducts: %s"
	id := converter_value.StringToInt(h.Ctx.Input.Param("id"))

	result, err := h.ProductUseCase.GetProductById(h.Ctx.Request.Context(), uint(id))

	if err != nil {
		h.log.Error(log,err.Error())
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
