package main

import (
	"clean-architecture-beego/database"
	"clean-architecture-beego/internal/domain"
	productHandler "clean-architecture-beego/internal/product/delivery/http"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	beego "github.com/beego/beego/v2/server/web"
	//_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	//_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	//_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	//_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"time"
)

func main() {

	//initialization database
	db := database.DB()

	if err := db.AutoMigrate(&domain.Product{}, &domain.Customer{}, &domain.Order{}); err != nil {
		panic(err)
	}

	//global timeout
	timeout, err := beego.AppConfig.Int("timeout")
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

	productRepository := productRepo.NewProductRepository(nil)
	productUseCase := productUcase.NewProductUseCase(timeoutContext,productRepository)
	productHandler.NewProductHandler(productUseCase)

	beego.Run()
}
