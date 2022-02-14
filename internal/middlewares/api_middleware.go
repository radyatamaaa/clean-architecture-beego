package middlewares

import (
	"clean-architecture-beego/pkg/logger"
	beego "github.com/beego/beego/v2/server/web"
	contextBeego "github.com/beego/beego/v2/server/web/context"
	"github.com/twinj/uuid"
	"io/ioutil"
	"net/http"
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
		requestId := uuid.NewV4().String()
		logging := logger.LoggingObj{
			RequestId: requestId,
		}


		ctx.Input.SetData("REQUEST_ID",requestId)
		ctx.Input.SetData("LOG",logging)

		next(ctx)

		requestbody,_ := ioutil.ReadAll(ctx.Request.Body)
		defer ctx.Request.Body.Close()

		var responsebody []byte
		ctx.ResponseWriter.Write(responsebody)
		logging = ctx.Input.GetData("LOG").(logger.LoggingObj)
		logging.Data.Request = string(requestbody)
		logging.Data.HttpCode = ctx.ResponseWriter.Status
		logging.Data.Method = ctx.Request.Method
		logging.Data.Response = string(responsebody)
		logging.Host = ctx.Request.Host
		logging.PathFile = ctx.Request.URL.String()
		if ctx.Input.GetData("FEATURE") != nil{
			logging.Feature = ctx.Input.GetData("FEATURE").(string)
		}
		if logging.Data.HttpCode != http.StatusOK {
			if ctx.Input.GetData("ERROR_MESSAGE") != nil{
				logging.Message = ctx.Input.GetData("ERROR_MESSAGE").(string)
			}
			m.Log.Error(logging)
			return
		}

		m.Log.Info(logging)
	}
}

