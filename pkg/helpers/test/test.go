package test

import (
	"clean-architecture-beego/pkg/helpers/converter_value"
	//"clean-architecture-beego/pkg/helpers/converter_value"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"reflect"
)

func NewMockEnv() error {

	dir, _ := os.Getwd()
	dir = filepath.Dir(dir)
	dir = filepath.Dir(dir) + "/"
	err := godotenv.Load(dir + ".env")
	if err != nil {
		return err
	}
	return nil
}

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func GetValueAndColumnStructToDriverValue(value interface{}) ([]driver.Value, []string) {
	var result []driver.Value

	//column
	j, _ := json.Marshal(value)
	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, e := range c {
		k[i] = s
		v, _ := e.MarshalJSON()
		var val driver.Value
		err := json.Unmarshal(v, &val)
		if err != nil {
			panic(err)
		}

		//dv := driver.Value(v.Interface())
		if reflect.TypeOf(val).String() == "string" {
			valueString := reflect.ValueOf(val).String()
			convertDate := converter_value.StringToDateTimeNullable(valueString)
			if convertDate.IsZero() == false {
				val = convertDate
			}
		}

		result = append(result, val)
		i++
	}

	return result, k

}
