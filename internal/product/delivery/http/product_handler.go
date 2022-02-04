package http

import (
	"clean-architecture-beego/pkg/helpers/api"
	"context"
	"clean-architecture-beego/internal/domain"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type ProductHandler struct {
	beego.Controller
	ProductUseCase domain.ProductUseCase
}

func NewProductHandler(useCase domain.ProductUseCase) {
	handler := &ProductHandler{
		ProductUseCase: useCase,
	}

	beego.Router("/api/v1/products", handler, "get:GetProducts")
	beego.Router("/api/v1/product/:id", handler, "get:GetProductByID")
	beego.Router("/api/v1/product", handler, "post:StoreProduct")
	beego.Router("/api/v1/product", handler, "put:UpdateProduct")
	beego.Router("/api/v1/product/:id", handler, "delete:DeleteProduct")
}

func (h *ProductHandler) GetProducts() {
	response := new(api.Response)
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
		response.MappingResponseError(api.GetStatusCode(err), err.Error() , err)
		h.Data["json"] = response
		if err := h.ServeJSON(); err != nil {
			return
		}
		return
	}
	response.MappingResponseSuccess("success", result)
	h.Data["json"] = response
	if err := h.ServeJSON(); err != nil {
		return
	}
	return
}

func (h *ProductHandler) StoreProduct() {

}

func (h *ProductHandler) UpdateProduct() {

}

func (h *ProductHandler) DeleteProduct() {

}

func (h *ProductHandler) GetProductByID() {

}
