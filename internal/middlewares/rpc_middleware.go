package middlewares

import (
	"clean-architecture-beego/pkg/jwt"
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JwtConfigRpc struct {
	Skipper func(ctx context.Context) bool
	AuthJwt jwt.JWT
}

func NewAuthFunc(ignoreMethod []string,jwtAuth jwt.JWT) *JwtConfigRpc {
	return &JwtConfigRpc{Skipper: func(ctx context.Context) bool {
		method, _ := grpc.Method(ctx)
		for _, imethod := range ignoreMethod {
			if method == imethod {
				return true
			}
		}
		return false
	},
	AuthJwt: jwtAuth}
}

func(c *JwtConfigRpc) AuthFunc(ctx context.Context) (context.Context, error) {
	if c.Skipper(ctx) {
		return ctx,nil
	}
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	tokenContext, err := c.parseToken(token,ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	return tokenContext, nil
}

func(c *JwtConfigRpc) parseToken(token string,ctx context.Context) (context.Context, error) {
	ctx ,err := c.AuthJwt.MiddlewareRPCAuth(ctx,token)
	if err != nil{
		return nil, err
	}
	return ctx, nil
}