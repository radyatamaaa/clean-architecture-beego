package main

import (
	"clean-architecture-beego/internal/middlewares"
	productGrpc "clean-architecture-beego/internal/product/delivery/grpc"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	"clean-architecture-beego/pkg/database"
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

	timeoutContext := time.Duration(requestTimeout) * time.Second

	if listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", grpcHost, httpPortGrpc)); err != nil {
		logs.Critical("Could not listen @ %v :: %v", httpPortGrpc, err)
	} else {
		var ignoreMethod = []string{"/product.ProductService/GetProductByID"}
		grpcServer := grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(middlewares.NewAuthFunc(ignoreMethod).AuthFunc),
			grpc_recovery.StreamServerInterceptor(),
		)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_auth.UnaryServerInterceptor(middlewares.NewAuthFunc(ignoreMethod).AuthFunc),
				grpc_recovery.UnaryServerInterceptor(),
			)),
			)

		//register Services
		productRepository := productRepo.NewProductRepository(db)
		productUseCase := productUcase.NewProductUseCase(timeoutContext, productRepository)
		productService := productGrpc.NewProductService(productUseCase)
		productGrpc.RegisterProductServiceServer(grpcServer, productService)

		//grpc listen and serve
		err = grpcServer.Serve(listen)
		if err != nil {
			logs.Critical("Failed to start gRPC Server :: %v", err)
		}
		logs.Info("service Running @ : " + httpPortGrpc)
	}
}
