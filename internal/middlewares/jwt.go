package middlewares

import (
	"clean-architecture-beego/pkg/helpers/response"
	"clean-architecture-beego/pkg/jwt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"strings"
)

type JwtConfig struct {
	Skipper func(ctx *context.Context) bool
}

func NewJwtMiddleware() *JwtConfig {
	return &JwtConfig{Skipper: func(ctx *context.Context) bool {
		if strings.EqualFold(ctx.Request.URL.Path, "/token") {
			return true
		}
		if strings.EqualFold(ctx.Request.URL.Path, "/register") {
			return true
		}
		return false
	}}
}

func (r *JwtConfig) JwtMiddleware(jwtAuth jwt.JWT) beego.FilterChain {
	return func(next beego.FilterFunc) beego.FilterFunc {
		return func(ctx *context.Context) {
			if r.Skipper(ctx) {
				next(ctx)
			}
			if r, err := jwtAuth.Middleware(ctx.Request); err != nil {
				ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
				switch {
				case jwt.IsInvalidToken(err):
					ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
					ctx.Output.JSON(response.ApiResponse{
						Code:    "INVALID_TOKEN",
						Message: "token is invalid",
					}, beego.BConfig.RunMode != "prod", false)
					return
				case jwt.IsExpiredToken(err):
					ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
					ctx.Output.JSON(response.ApiResponse{
						Code:    "EXPIRED_TOKEN",
						Message: "token is expired",
					}, beego.BConfig.RunMode != "prod", false)
					return
				case jwt.IsMissingToken(err):
					ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
					ctx.Output.JSON(response.ApiResponse{
						Code:    "MISSING_TOKEN",
						Message: "token is missing",
					}, beego.BConfig.RunMode != "prod", false)
					return
				case jwt.IsAuthElsewhere(err):
					ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
					ctx.Output.JSON(response.ApiResponse{
						Code:    "AUTH_ELSEWHERE",
						Message: "auth elsewhere",
					}, beego.BConfig.RunMode != "prod", false)
					return
				default:
					ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
					ctx.Output.JSON(response.ApiResponse{
						Code:    response.UnauthorizedError,
						Message: response.ErrUnAuthorize.Error(),
					}, beego.BConfig.RunMode != "prod", false)
					return
				}
			} else {
				ctx.Request = r
				next(ctx)
			}
		}
	}
}
