package routers

import (
	productGrpc "clean-architecture-beego/internal/product/delivery/grpc"
	productHandler "clean-architecture-beego/internal/product/delivery/http"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"strconv"
	"time"
)

func InitializeRouter(db *gorm.DB, durationTimeout time.Duration, httpPortGrpc int) {

	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
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

	productRepository := productRepo.NewProductRepository(db)
	productUseCase := productUcase.NewProductUseCase(durationTimeout, productRepository)
	pHandler := productHandler.NewProductHandler(productUseCase)


	ns :=
		web.NewNamespace("/api/v1",
			web.NSNamespace("/product",
				web.NSInclude(
					&pHandler,
				),
			),
		)
	web.AddNamespace(ns)

	if listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", httpPortGrpc)); err != nil {
		log.Fatalf("Could not listen @ %v :: %v", httpPortGrpc, err)
	} else {
		grpcServer := grpc.NewServer()

		//register Services
		productService := productGrpc.NewProductService(productUseCase)
		productGrpc.RegisterProductServiceServer(grpcServer, productService)

		//grpc listen and serve
		err = grpcServer.Serve(listen)
		if err != nil {
			log.Fatalf("Failed to start gRPC Server :: %v", err)
		}
		log.Println("service Running @ : " + strconv.Itoa(httpPortGrpc))
	}
}
