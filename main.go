package main

import (
	"clean-architecture-beego/database"
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/routers"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

func main() {

	// default vars
	var (
		requestTimeout = 30
		httpPortGrpc   = 9090
	)

	//initialization database
	db := database.DB()

	if err := db.AutoMigrate(&domain.Product{}, &domain.Customer{}, &domain.Order{}); err != nil {
		panic(err)
	}

	//global timeout
	if timeout, err := beego.AppConfig.Int("timeout"); err == nil {
		requestTimeout = timeout
	}
	if port, err := beego.AppConfig.Int("httpportGRPC"); err == nil {
		httpPortGrpc = port
	}

	timeoutContext := time.Duration(requestTimeout) * time.Second

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// init router
	routers.InitializeRouter(db, timeoutContext, httpPortGrpc)

	beego.Run()

}
