package middlewares

import (
	"clean-architecture-beego/pkg/jwt"
	"clean-architecture-beego/pkg/logger"
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type RpcMiddleware struct {
	Skipper func(ctx context.Context) bool
	AuthJwt jwt.JWT
	Log logger.Logger
}

func NewRpcMiddleware(ignoreMethod []string,jwtAuth jwt.JWT,logger logger.Logger) *RpcMiddleware {
	return &RpcMiddleware{Skipper: func(ctx context.Context) bool {
		method, _ := grpc.Method(ctx)
		for _, imethod := range ignoreMethod {
			if method == imethod {
				return true
			}
		}
		return false
	},
	AuthJwt: jwtAuth,
	Log: logger}
}
func(c *RpcMiddleware) LoggerStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	c.Log.Info("Accepted")

	ctx := stream.Context()
	status := codes.OK
	statusDesc := codes.OK.String()
	method, _ := grpc.Method(ctx)
	err = handler(srv, stream)
	if err != nil{
		status = grpc.Code(err)
		statusDesc = grpc.Code(err).String()
	}

	c.Log.Info("[" + strconv.Itoa(int(status)) + "] " + "[" + statusDesc + "] " + method + " ")

	c.Log.Info("Closing")

	return nil
}
func(c *RpcMiddleware) LoggerUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	c.Log.Info("Accepted")

	status := codes.OK
	statusDesc := codes.OK.String()
	method, _ := grpc.Method(ctx)
	_,err := handler(ctx, req)
	if err != nil{
		status = grpc.Code(err)
		statusDesc = grpc.Code(err).String()
	}

	c.Log.Info("[" + strconv.Itoa(int(status)) + "] " + "[" + statusDesc + "] " + method + " ")

	c.Log.Info("Closing")

	return nil,nil
}
func(c *RpcMiddleware) AuthFunc(ctx context.Context) (context.Context, error) {
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

func(c *RpcMiddleware) parseToken(token string,ctx context.Context) (context.Context, error) {
	ctx ,err := c.AuthJwt.MiddlewareRPCAuth(ctx,token)
	if err != nil{
		return nil, err
	}
	return ctx, nil
}