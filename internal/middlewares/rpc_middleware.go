package middlewares

import (
	"clean-architecture-beego/pkg/jwt"
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	tokenContext, err := parseToken(token,ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	//grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))

	// WARNING: in production define your own type to avoid context collisions
	//newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)

	return tokenContext, nil
}

func parseToken(token string,ctx context.Context) (context.Context, error) {
	auth, err := jwt.NewJwt(&jwt.Options{
		Issuer:      "backend",
		SignMethod:  jwt.HS256,
		SecretKey:   "secret",
		ExpiredTime: 1500,
		Locations:   "header:Authorization",
		IdentityKey: "uid",
	})
	if err != nil{
		panic(err)
	}

	ctx ,err = auth.MiddlewareRPCAuth(ctx,token)
	if err != nil{
		return nil, err
	}
	return ctx, nil
}

func userClaimFromToken(struct{}) string {
	return "foobar"
}