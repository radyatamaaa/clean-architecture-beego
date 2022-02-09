package main

import (
	"clean-architecture-beego/internal/domain"
	productHandler "clean-architecture-beego/internal/product/delivery/http"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	userHandler "clean-architecture-beego/internal/user/delivery/http"
	userRepo "clean-architecture-beego/internal/user/repository"
	userUcase "clean-architecture-beego/internal/user/usecase"
	"clean-architecture-beego/pkg/database"
	"clean-architecture-beego/pkg/helpers/response"
	"clean-architecture-beego/pkg/jwt"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
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
	if err := db.AutoMigrate(&domain.Product{}, &domain.Customer{}, &domain.Order{}, &domain.User{}); err != nil {
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

	// middleware cors
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:    []string{http.MethodGet, http.MethodPost},
		AllowAllOrigins: true,
	}))

	// jwt middleware
	if auth, err := jwt.NewJwt(&jwt.Options{
		Issuer:      "backend",
		SignMethod:  jwt.HS256,
		SecretKey:   "secret",
		ExpiredTime: 1500,
		Locations:   "header:Authorization",
		IdentityKey: "uid",
	}); err == nil {
		if bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`); err != nil {
			panic(err)
		} else {
			auth.SetAdapter(bm)

			// user handler
			userRepository := userRepo.NewUserRepository(db)
			userUseCase := userUcase.NewUserUseCase(timeoutContext, userRepository)
			userHandler.NewUserHandler(userUseCase, auth)

			beego.InsertFilter("/api/*", beego.BeforeRouter, func(ctx *context.Context) {
				if r, err := auth.Middleware(ctx.Request); err != nil {
					ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
					switch {
					case jwt.IsInvalidToken(err):
						ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
						ctx.Output.JSON(response.ApiResponse{
							Code:    "INVALID_TOKEN",
							Message: "token is invalid",
						}, beego.BConfig.RunMode != "prod", false)
						return
					case jwt.IsExpiredToken(err):
						ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
						ctx.Output.JSON(response.ApiResponse{
							Code:    "EXPIRED_TOKEN",
							Message: "token is expired",
						}, beego.BConfig.RunMode != "prod", false)
						return
					case jwt.IsMissingToken(err):
						ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
						ctx.Output.JSON(response.ApiResponse{
							Code:    "MISSING_TOKEN",
							Message: "token is missing",
						}, beego.BConfig.RunMode != "prod", false)
						return
					case jwt.IsAuthElsewhere(err):
						ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
						ctx.Output.JSON(response.ApiResponse{
							Code:    "AUTH_ELSE_WHERE",
							Message: "auth else where",
						}, beego.BConfig.RunMode != "prod", false)
						return
					default:
						ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
						ctx.Output.JSON(response.ApiResponse{
							Code:    response.UnauthorizedError,
							Message: response.ErrUnAuthorize.Error(),
						}, beego.BConfig.RunMode != "prod", false)
						return
					}
				} else {
					ctx.Request = r
				}
			})
		}
	} else {
		panic(err)
	}

	// default error handler
	beego.ErrorController(&response.ErrorController{})

	// product handler
	productRepository := productRepo.NewProductRepository(db)
	productUseCase := productUcase.NewProductUseCase(timeoutContext, productRepository)
	productHandler.NewProductHandler(productUseCase)

	beego.Run()
}
