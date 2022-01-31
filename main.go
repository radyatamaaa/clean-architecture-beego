package main

import (
	"clean-architecture-beego/database"
	"clean-architecture-beego/internal/domain"
	productHandler "clean-architecture-beego/internal/product/delivery/http"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {

	// initialization database
	db := database.DB()

	if err := db.AutoMigrate(&domain.Product{}, &domain.Customer{}, &domain.Order{}); err != nil {
		panic(err)
	}
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

	productRepository := productRepo.NewProductRepository(db)
	productUseCase := productUcase.NewProductUseCase(productRepository)
	productHandler.NewProductHandler(productUseCase)

	beego.Run()
}
