package main

import (
	productGrpc "clean-architecture-beego/internal/product/delivery/grpc"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	"clean-architecture-beego/pkg/database"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
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
		grpcServer := grpc.NewServer()

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
