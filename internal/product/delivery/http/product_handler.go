package http

import (
	"clean-architecture-beego/internal/domain"
	"context"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type ProductHandler struct {
	beego.Controller
	ProductUseCase domain.ProductUseCase
}

func NewProductHandler(useCase domain.ProductUseCase) ProductHandler {
	return ProductHandler{
		ProductUseCase: useCase,
	}
}

func (h *ProductHandler) URLMapping() {
	h.Mapping("GetProducts", h.GetProducts)
	h.Mapping("StoreProduct", h.StoreProduct)
	h.Mapping("UpdateProduct", h.UpdateProduct)
	h.Mapping("DeleteProduct", h.Delete)
	h.Mapping("GetProductByID", h.GetProductByID)

}

// GetProducts get all products
// @router / [get]
func (h *ProductHandler) GetProducts() {

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
	result, err := h.ProductUseCase.GetProducts(ctx,limit, offset)

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

// StoreProduct save product
// @router / [post]
func (h *ProductHandler) StoreProduct() {

}

// UpdateProduct update products
// @router / [put]
func (h *ProductHandler) UpdateProduct() {

}

// DeleteProduct get delete products
// @router /:id [delete]
func (h *ProductHandler) DeleteProduct() {

}

// GetProductByID product by id
// @router /:id [get]
func (h *ProductHandler) GetProductByID() {

}
