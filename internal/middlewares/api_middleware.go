package middlewares

import (
	"clean-architecture-beego/pkg/logger"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"

	"github.com/beego/beego/v2/server/web/context"
)

type Middleware struct {
	Log logger.Logger
}
func NewMiddleware(logger logger.Logger) *Middleware {
	return &Middleware{
		Log: logger,
	}
}
func (m *Middleware) Logger(next beego.FilterFunc) beego.FilterFunc {
	return func(ctx *context.Context) {
		m.Log.Info("Accepted")

		next(ctx)

		m.Log.Info("[" + strconv.Itoa(ctx.ResponseWriter.Status) + "] " + "[" + ctx.Request.Method + "] " + ctx.Request.Host + ctx.Request.URL.String())

		m.Log.Info("Closing")
	}
}

