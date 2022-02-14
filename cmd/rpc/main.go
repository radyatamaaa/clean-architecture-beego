package main

import (
	"clean-architecture-beego/internal/middlewares"
	productGrpc "clean-architecture-beego/internal/product/delivery/grpc"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	"clean-architecture-beego/pkg/database"
	"clean-architecture-beego/pkg/jwt"
	"clean-architecture-beego/pkg/logger"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"net"
	"time"
)

// sample grpc server
func main() {

	// default vars
	var (
		requestTimeout = 3
		httpPortGrpc   = "9090"
		grpcHost       = "127.0.0.1"
	)

	if err := beego.LoadAppConfig("ini", "./conf/app.conf"); err != nil {
		panic(err)
	}

	db := database.DB()

	// global timeout
	if timeout, err := beego.AppConfig.Int("timeout"); err == nil {
		requestTimeout = timeout
	}

	// rpc host
	if host, err := beego.AppConfig.String("hostRpc"); err == nil {
		grpcHost = host
	}

	// rpc port
	if port, err := beego.AppConfig.String("portRpc"); err == nil {
		httpPortGrpc = port
	}

	appname ,err := beego.AppConfig.String("appname")
	if err != nil {
		panic(err)
	}

	app ,err := beego.AppConfig.String("app")
	if err != nil {
		panic(err)
	}

	version ,err := beego.AppConfig.String("version")
	if err != nil {
		panic(err)
	}
	
	timeoutContext := time.Duration(requestTimeout) * time.Second

	// logger
	l := logger.NewStdOutLogger(30,"all","Local",true,version,app,appname)

	if listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", grpcHost, httpPortGrpc)); err != nil {
		logs.Critical("Could not listen @ %v :: %v", httpPortGrpc, err)
	} else {
		ignoreMethod := []string{"/product.ProductService/GetProductByID"}
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
		rpcMiddlewares := middlewares.NewRpcMiddleware(ignoreMethod,auth,l)

		grpcServer := grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				rpcMiddlewares.LoggerStreamInterceptor,
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(rpcMiddlewares.AuthFunc),
			grpc_recovery.StreamServerInterceptor(),


		)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				rpcMiddlewares.LoggerUnaryServerInterceptor,
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_auth.UnaryServerInterceptor(rpcMiddlewares.AuthFunc),
				grpc_recovery.UnaryServerInterceptor(),
			)),

			)

		//register Services
		productRepository := productRepo.NewProductRepository(db,l)
		productUseCase := productUcase.NewProductUseCase(timeoutContext, productRepository,l)
		productService := productGrpc.NewProductService(productUseCase,l)
		productGrpc.RegisterProductServiceServer(grpcServer, productService)

		//grpc listen and serve
		err = grpcServer.Serve(listen)
		if err != nil {
			logs.Critical("Failed to start gRPC Server :: %v", err)
		}
		logs.Info("service Running @ : " + httpPortGrpc)
	}
}
