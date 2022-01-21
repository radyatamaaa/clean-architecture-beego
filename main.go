package main

import (
	"clean-architecture-beego/helper/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"

	_sampleModuleHttpHandler "clean-architecture-beego/sample_module/delivery/http"
	_sampleModuleUsecase "clean-architecture-beego/sample_module/usecase"
	beego "github.com/beego/beego/v2/server/web"

	_sampleModuleRepo "clean-architecture-beego/sample_module/repository"

	_ "clean-architecture-beego/swagger"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

// @title Echo Swagger clean-architecture-beego API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//dbConn := dbconn.DB()
	//var dbConn *gorm.DB = nil
	l := logger.L
	//appPort := os.Getenv("PORT")
	timeout, _ := strconv.Atoi(os.Getenv("APP_TIMEOUT"))
	//env := os.Getenv("APP_ENV")

	sampleModuleRepo := _sampleModuleRepo.NewSampleModuleRepository(nil, l)

	timeoutContext := time.Duration(timeout) * time.Second

	sampleModuleUsecase := _sampleModuleUsecase.NewSampleModuleUsecase(sampleModuleRepo, l, timeoutContext)

	_sampleModuleHttpHandler.NewsampleModuleHandler( sampleModuleUsecase, l)

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


	beego.Run()
	//log.Fatal(e.Start(":" + appPort))
}
