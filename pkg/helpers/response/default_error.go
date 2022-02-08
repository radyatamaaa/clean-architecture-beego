package response

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type ErrorController struct {
	web.Controller
	ApiResponse
}

func (c *ErrorController) Error404() {
	c.ErrorResponse(c.Ctx, http.StatusNotFound, ResourceNotFoundError, nil)
	return
}
