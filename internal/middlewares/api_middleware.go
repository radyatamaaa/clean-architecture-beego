package middlewares

import (
	"clean-architecture-beego/pkg/logger"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	context "context"
	contextBeego "github.com/beego/beego/v2/server/web/context"
	"github.com/twinj/uuid"
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
	return func(ctx *contextBeego.Context) {
		m.Log.Info("Accepted")

		requestId := uuid.NewV4().String()
		c := context.WithValue(ctx.Request.Context(), "REQUEST_ID", requestId)
		ctx.Request.WithContext(c)

		m.Log.Info("Request-Id : " + requestId)

		next(ctx)

		m.Log.Info("[" + strconv.Itoa(ctx.ResponseWriter.Status) + "] " + "[" + ctx.Request.Method + "] " + ctx.Request.Host + ctx.Request.URL.String())

		m.Log.Info("Closing")
	}
}

