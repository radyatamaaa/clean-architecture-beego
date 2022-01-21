package http

import (
	"context"
	"clean-architecture-beego/helper"
	"clean-architecture-beego/helper/logger"
	models2 "clean-architecture-beego/helper/models"
	"clean-architecture-beego/sample_module"
	_ "clean-architecture-beego/swagger/models"
	"net/http"
	beego "github.com/beego/beego/v2/server/web"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// sampleModuleHandler  represent the http handler for country
type SampleModuleHandler struct {
	beego.Controller
	SampleModuleUsecase sample_module.Usecase
	Log                 logger.Logger
}

func (c *SampleModuleHandler) URLMapping() {
	c.Mapping("List", c.Get)
	c.Mapping("Documentation", c.Get)
}

// NewsampleModuleHandler will initialize the countrys/ resources endpoint
func NewsampleModuleHandler(us sample_module.Usecase, log logger.Logger) {
	handler := &SampleModuleHandler{
		Log:                 log,
		SampleModuleUsecase: us,
	}

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/sample-module",
			beego.NSInclude(
				handler,
			),
		),
	)
	beego.AddNamespace(ns)

}
func (a *SampleModuleHandler) Documentation()  {
	http.Redirect(a.Ctx.ResponseWriter, a.Ctx.Request, "/swagger/index.html", 301)
	return
}

// List godoc
// @Summary List.
// @Description List.
// @Tags sampleModule
// @Accept  json
// @Produce  json
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Success 200 {object} models.SwaggerListSuccess
// @Failure 404 {object} models.SwaggerErrorNotFound
// @Failure 409 {object} models.SwaggerErrorConflict
// @Failure 401 {object} models.SwaggerErrorUnAuthorize
// @Failure 400 {object} models.SwaggerErrorBadParamInput
// @Failure 500 {object} models.SwaggerOtherInternalServerError
// @Failure 400 {object} models.SwaggerErrorInvalidMethod
// @Failure 405 {object} models.SwaggerErrorMethodNotAllowed
// @Router /api/sample-module [get]
func (a *SampleModuleHandler) List()  {
	response := new(models2.Response)

	qpage := a.Ctx.Input.Query("page")
	qperPage := a.Ctx.Input.Query("limit")

	ctx := a.Ctx.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	limit, page, offset := helper.Pagination(qpage, qperPage)
	result, message, err := a.SampleModuleUsecase.GetList(ctx, page, limit, offset)
	if err != nil {

		a.Log.Error("sample_module.delivery.http.SampleModuleHandler.List: %s", err.Error())
		response.MappingResponseError(helper.GetStatusCode(err), message, err.Error())
		a.Data["json"] = response
		a.Ctx.Output.SetStatus(response.StatusCode)
		a.ServeJSON()
		return
	}
	response.MappingResponseSuccess(message, result)
	a.Data["json"] = response
	a.Ctx.Output.SetStatus(response.StatusCode)
	a.ServeJSON()
	return
}
