package main

import (
	"clean-architecture-beego/internal/domain"
	productHandler "clean-architecture-beego/internal/product/delivery/http"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	"clean-architecture-beego/pkg/database"
	"clean-architecture-beego/pkg/helpers/response"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"net/http"
	"time"
)
// @title BEE SAMPLE API
// @version v1
// @contact.name Kredit Plus
// @contact.url https://kreditplus.com
// @contact.email support@kreditplus.com
// @description api "sample using beego framework"
// @termsOfService https://dev-kreditmu.kreditplus.com/terms

// @BasePath /api
// @query.collection.format multi
func main() {

	// default vars
	var (
		requestTimeout = 3
	)

	if err := beego.LoadAppConfig("ini", "./conf/app.conf"); err != nil {
		panic(err)
	}

	// initialization database
	db := database.DB()

	// migrate database
	if err := db.AutoMigrate(&domain.Product{}, &domain.Customer{}, &domain.Order{}); err != nil {
		panic(err)
	}

	// global timeout
	if timeout, err := beego.AppConfig.Int("timeout"); err == nil {
		requestTimeout = timeout
	}

	timeoutContext := time.Duration(requestTimeout) * time.Second

	// swagger config
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:    []string{http.MethodGet, http.MethodPost},
		AllowAllOrigins: true,
	}))

	// default error handler
	beego.ErrorController(&response.ErrorController{})

	// product handler
	productRepository := productRepo.NewProductRepository(db)
	productUseCase := productUcase.NewProductUseCase(timeoutContext, productRepository)
	productHandler.NewProductHandler(productUseCase)

	beego.Run()
}
