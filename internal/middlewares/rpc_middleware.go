package middlewares

import (
	"clean-architecture-beego/pkg/jwt"
	"clean-architecture-beego/pkg/logger"
	"context"
	"encoding/json"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/twinj/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	ctx := stream.Context()

	requestId := uuid.NewV4().String()
	logging := logger.LoggingObj{
		RequestId: requestId,
	}

	ctx = context.WithValue(ctx, "REQUEST_ID", requestId)
	ctx = context.WithValue(ctx, "LOG", logging)

	status := codes.OK
	statusDesc := codes.OK.String()
	method, _ := grpc.Method(ctx)
	err = handler(srv, stream)
	if err != nil{
		status = grpc.Code(err)
		statusDesc = grpc.Code(err).String()
	}
	requestbody,_ := json.Marshal(srv)
	responsebody,_ := json.Marshal(stream)

	logging = ctx.Value("LOG").(logger.LoggingObj)
	logging.Data.Request = string(requestbody)
	logging.Data.HttpCode = int(status)
	logging.Data.HttpCodeDesc = statusDesc
	logging.Data.Method = ""
	logging.Data.Response = string(responsebody)
	logging.Host = ""
	logging.PathFile = method
	if ctx.Value("FEATURE") != nil{
		logging.Feature = ctx.Value("FEATURE").(string)
	}
	if logging.Data.HttpCode != int(codes.OK) {
		if ctx.Value("ERROR_MESSAGE") != nil{
			logging.Message = ctx.Value("ERROR_MESSAGE").(string)
		}
		c.Log.Error(logging)
		return nil
	}

	c.Log.Info(logging)

	return nil
}
func(c *RpcMiddleware) LoggerUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {


	requestId := uuid.NewV4().String()
	logging := logger.LoggingObj{
		RequestId: requestId,
	}

	ctx = context.WithValue(ctx, "REQUEST_ID", requestId)
	ctx = context.WithValue(ctx, "LOG", logging)

	status := codes.OK
	statusDesc := codes.OK.String()
	method, _ := grpc.Method(ctx)
	res,err := handler(ctx, req)
	if err != nil{
		status = grpc.Code(err)
		statusDesc = grpc.Code(err).String()
	}
	requestbody,_ := json.Marshal(req)
	responsebody,_ := json.Marshal(res)

	logging = ctx.Value("LOG").(logger.LoggingObj)
	logging.Data.Request = string(requestbody)
	logging.Data.HttpCode = int(status)
	logging.Data.HttpCodeDesc = statusDesc
	logging.Data.Method = ""
	logging.Data.Response = string(responsebody)
	logging.Host = ""
	logging.PathFile = method
	if ctx.Value("FEATURE") != nil{
		logging.Feature = ctx.Value("FEATURE").(string)
	}
	if logging.Data.HttpCode != int(codes.OK) {
		if ctx.Value("ERROR_MESSAGE") != nil{
			logging.Message = ctx.Value("ERROR_MESSAGE").(string)
		}
		c.Log.Error(logging)
		return nil,nil
	}

	c.Log.Info(logging)

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