package dbconn

import (
	"os"
	"sync"

	_ "github.com/apache/calcite-avatica-go/v5"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// singleton instance of database connection.
var dbInstance *gorm.DB
var dbOnce sync.Once

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
	gormDB, err := gorm.Open(
		sqlserver.Open(os.Getenv("DB_CONNECTION_URL")),
		&gorm.Config{SkipDefaultTransaction: true},
	)

	if err != nil {
		panic("dbconn openDB PortalBRIBrain: cannot open database")
	}
	dbInstance = gormDB
	db, err := dbInstance.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(200)
}
