package database

import (
	"fmt"
	_ "github.com/apache/calcite-avatica-go/v5"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

// singleton instance of database connection.
var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

// DB creates a new instance of gorm.DB if a connection is not established.
// return singleton instance.
func DB() *gorm.DB {
	if dbInstance == nil {
		dbOnce.Do(openDB)
	}
	return dbInstance
}

// openDB initialize gorm DB.
func openDB() {
	dsn, err := beego.AppConfig.String("dbConnUrl")
	if err != nil {
		panic(err)
	}
	gormDB, err := gorm.Open(
		sqlserver.Open(dsn),
		&gorm.Config{SkipDefaultTransaction: true,Logger:  logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		panic("cannot open database.")
	}
	dbInstance = gormDB
	db, err := dbInstance.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(200)

	fmt.Println("success connect to db")
}
