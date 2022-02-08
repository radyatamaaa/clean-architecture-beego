package test

import (
	"clean-architecture-beego/pkg/helpers/converter_value"
	beego "github.com/beego/beego/v2/server/web"
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"

	//"clean-architecture-beego/pkg/helpers/converter_value"
	"database/sql/driver"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"path/filepath"
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

func PrepareHandler(t *testing.T,handler * beego.Controller,request *http.Request,response http.ResponseWriter) {
	err := request.ParseForm()
	assert.NoError(t, err)

	handler.Ctx = &beegoContext.Context{
		Request:        request,
		ResponseWriter: &beegoContext.Response{
			ResponseWriter: response,
			Started:        false,
			Status:         0,
			Elapsed:        0,
		},
	}
	body, _ := ioutil.ReadAll(handler.Ctx.Request.Body)
	handler.Ctx.Input = &beegoContext.BeegoInput{
		Context:       handler.Ctx,
		CruSession:    nil,
		RequestBody:   body,
		RunMethod:     "",
		RunController: nil,
	}
	handler.Ctx.Output = &beegoContext.BeegoOutput{
		Context:    handler.Ctx,
	}
	handler.Data = map[interface{}]interface{} {}

}
