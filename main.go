package main

import (
	"clean-architecture-beego/database"
	customerGrpc "clean-architecture-beego/internal/customer/delivery/grpc"
	customerHandler "clean-architecture-beego/internal/customer/delivery/http"
	customerRepo "clean-architecture-beego/internal/customer/repository"
	customerUcase "clean-architecture-beego/internal/customer/usecase"
	"clean-architecture-beego/internal/domain"
	productGrpc "clean-architecture-beego/internal/product/delivery/grpc"
	productHandler "clean-architecture-beego/internal/product/delivery/http"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	"log"
	"net"

	beego "github.com/beego/beego/v2/server/web"
	"google.golang.org/grpc"

	//_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	//_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	//_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	//_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	"time"

	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {

	//initialization database
	db := database.DB()

	if err := db.AutoMigrate(&domain.Product{}, &domain.Customer{}, &domain.Order{}); err != nil {
		panic(err)
	}

	//global timeout
	timeout, err := beego.AppConfig.Int("timeout")
	httpportGRPC, err := beego.AppConfig.String("httpportGRPC")
	if err != nil {
		panic(err)
	}

	timeoutContext := time.Duration(timeout) * time.Second

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//Product
	productRepository := productRepo.NewProductRepository(db)
	productUseCase := productUcase.NewProductUseCase(timeoutContext, productRepository)
	productHandler.NewProductHandler(productUseCase)

	//Customer
	customerRepository := customerRepo.NewCustomerRepository(db)
	customerUseCase := customerUcase.NewCustomerUseCase(timeoutContext, customerRepository)
	customerHandler.NewCustomerHandler(customerUseCase)

	go func() {
		beego.Run()
	}()

	listen, err := net.Listen("tcp", ":"+httpportGRPC)
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", httpportGRPC, err)
	}
	log.Println("service Running @ : " + httpportGRPC)

	grpcserver := grpc.NewServer()

	//register Services
	//product
	productService := productGrpc.NewProductService(productUseCase)
	productGrpc.RegisterProductServiceServer(grpcserver, productService)

	//customer
	customerService := customerGrpc.NewCustomerService(customerUseCase)
	customerGrpc.RegisterCustomerServiceServer(grpcserver, customerService)

	//grpc listen and serve
	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}

}
