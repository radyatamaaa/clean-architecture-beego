package helper

import (
	"database/sql/driver"
	b64 "encoding/base64"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"math/rand"
	models2 "clean-architecture-beego/helper/models"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	Cache                 = cache.New(24*time.Hour, 24*time.Hour)
	DateTimeFormatDefault = "2006-01-02 15:04:05"
	DateFormatDefault = "2006-01-02"
)

type Emp []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SetCache(key string, emp interface{}, exp time.Duration) bool {
	Cache.Set(key, emp, exp)
	return true
}
func InArray(str string, list []string) bool {
	str = strings.ToLower(str)
	for _, v := range list {
		if strings.ToLower(v) == str {
			return true
		}
	}
	return false
}
func InArrayWithIndex(str string, list []string) (bool,*int) {
	str = strings.ToLower(str)
	for i, v := range list {
		if strings.ToLower(v) == str {
			return true,&i
		}
	}
	return false,nil
}
func GetCacheToken(key string) (string, bool) {
	var emp *time.Time
	var found bool
	data, found := Cache.Get(key)
	if found {
		emp = data.(*time.Time)
		return emp.String(), found
	}

	return "", found
}
func GetCache(key string) (string, bool) {
	var emp int
	var found bool
	data, found := Cache.Get(key)
	if found {
		emp = data.(int)
	}

	return strconv.Itoa(emp), found
}

func Pagination(qpage, qperPage string) (limit, page, offset int) {
	limit = 20
	page = 1
	offset = 0

	page, _ = strconv.Atoi(qpage)
	limit, _ = strconv.Atoi(qperPage)
	if page == 0 && limit == 0 {
		page = 1
		limit = 10
	}
	offset = (page - 1) * limit

	return
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case models2.ErrInternalServerError:
		return http.StatusInternalServerError
	case models2.ErrForbidden:
		return http.StatusForbidden
	case models2.ErrNotFound:
		return http.StatusNotFound
	case models2.PersonalNumberNotFound:
		return http.StatusUnauthorized
	case models2.LoginFailedMessageLockedUser:
		return http.StatusBadRequest
	case models2.ErrUnAuthorize:
		return http.StatusUnauthorized
	case models2.ErrConflict:
		return http.StatusConflict
	case models2.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func GetKeyJsonStruct(value interface{}) []string {

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
	for s, _ := range c {
		k[i] = s
		i++
	}

	return k
}

func GetValueStruct(value interface{}) []driver.Value {
	var result []driver.Value
	rv := reflect.ValueOf(value)
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)

		dv := driver.Value(fv.Interface())
		result = append(result, dv)
	}
	return result
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
			convertDate := StringToDateTimeNullable(valueString)
			if convertDate.IsZero() == false {
				val = convertDate
			}
		}

		result = append(result, val)
		i++
	}

	return result, k

}

func NowYmd() string {
	t := time.Now()
	timeFormated := t.Format(DateTimeFormatDefault)
	return timeFormated
}
func FloatToString(input_num float64) string {
	// to convert a float number to a string
	if input_num != 0 {
		return strconv.FormatFloat(input_num, 'f', 0, 64)
	} else {
		return "0"
	}
}

func StringNullableToFloat(value *string) float64 {
	if value != nil {
		res, _ := strconv.ParseFloat(*value, 64)
		return res
	}
	return 0
}
func StringToFloat(value string) float64 {
	if value != "" {
		res, _ := strconv.ParseFloat(value, 64)
		return res
	}
	return 0
}
func FloatNUllableToString(input_num *float64) string {
	// to convert a float number to a string
	if input_num != nil {
		return strconv.FormatFloat(*input_num, 'f', 0, 64)
	} else {
		return ""
	}
}
func FloatNUllableToFloat(value *float64) float64 {
	if value != nil {
		return *value
	}
	return 0
}
func FloatToFloatNullable(value float64) *float64 {
	return &value
}
func DateTimeToDateTimeNUllable(value time.Time) *time.Time {
	return &value
}
func DateTimeNullableToDateTime(value *time.Time) time.Time {
	if value == nil {
		return time.Time{}
	}
	return *value
}
func IntToIntNullable(value int) *int {
	return &value
}
func IntNullableToInt(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}
func StringToStringNullable(value string) *string {
	return &value
}
func ObjectToString(value interface{}) string {
	result, _ := json.Marshal(value)
	return string(result)
}
func StringNullableToString(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}
func IntNullableToStringNullable(value *int) *string {

	if value != nil {
		result := strconv.Itoa(*value)
		return &result
	}
	return nil
}
func IntNullableToString(value *int) string {

	if value != nil {
		result := strconv.Itoa(*value)
		return result
	}
	return "0"
}

func IntToString(value int) string {

	if value != 0 {
		result := strconv.Itoa(value)
		return result
	}
	return "0"
}
func StringToIntNullable(value string) *int {

	if value != "" {
		result, _ := strconv.Atoi(value)
		return &result
	}
	return nil
}
func Int64NullableToInt(value *int64) int {

	if value != nil {
		result := int(*value)
		return result
	}
	return 0
}
func StringToInt(value string) int {

	if value != "" {
		result, _ := strconv.Atoi(value)
		return result
	}
	return 0
}
func StringNullableToInt(value *string) int {

	if value != nil {
		result, _ := strconv.Atoi(*value)
		return result
	}
	return 0
}
func StringNullableToDateTimeNullable(value *string) *time.Time {
	if value != nil {
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02 15:04:05"
		date, _ = time.Parse(layoutFormat, *value)
		return &date
	}

	return nil
}

func DateTimeNullableToStringNullable(value *time.Time) *string {
	if value != nil {
		layoutFormat := "2006-01-02 15:04:05"
		date := value.Format(layoutFormat)
		return &date
	}

	return nil
}

func DateTimeToStringNullable(value time.Time) *string {
	layoutFormat := "2006-01-02 15:04:05"
	date := value.Format(layoutFormat)
	return &date
}
func DateTimeToStringWithFormat(value time.Time, format string) string {
	if !value.IsZero() {
		layoutFormat := format
		date := value.Format(layoutFormat)
		return date
	}

	return ""
}
func DateTimeNullableToStringNullableWithFormat(value *time.Time, format string) *string {
	if value != nil {
		layoutFormat := format
		date := value.Format(layoutFormat)
		return &date
	}

	return nil
}

func StringNullableToStringDefaultFormatDate(value *string) *string {
	if value != nil {
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02T15:04:05Z"
		date, _ = time.Parse(layoutFormat, *value)
		dateString := date.Format(DateTimeFormatDefault)
		return &dateString
	}

	return nil
}
func StringNullableToDateTime(value *string) time.Time {
	if value != nil {
		var layoutFormat string
		var date time.Time
		layoutFormat = "2006-01-02T15:04:05Z"
		date, err := time.Parse(layoutFormat, *value)
		if err != nil {
			return time.Time{}
		}
		return date
	}

	return time.Time{}
}
func StringToDateTimeNullable(value string) time.Time {
	if value != "" {
		var layoutFormat string
		var date time.Time
		layoutFormat = "2006-01-02T15:04:05.999999999Z07:00"
		date, err := time.Parse(layoutFormat, value)
		if err != nil {
			return time.Time{}
		}
		return date
	}

	return time.Time{}
}
func StringToDateWithFormat(value string,format string) time.Time {
	if value != "" {
		var layoutFormat string
		var date time.Time

		layoutFormat = format
		date, _ = time.Parse(layoutFormat, value)
		return date
	}

	return time.Time{}
}
func StringToDate(value string) time.Time {
	if value != "" {
		var layoutFormat string
		var date time.Time

		layoutFormat = DateFormatDefault
		date, _ = time.Parse(layoutFormat, value)
		return date
	}

	return time.Time{}
}
func StringNullableToDateNullable(value *string) *string {
	if value != nil {
		var layoutFormat string
		var date time.Time

		layoutFormat = "20060102"
		date, _ = time.Parse(layoutFormat, *value)
		dateString := date.Format("20060102")
		return &dateString
	}

	return nil
}
func ConvertIntBool(value *int) bool {
	if value != nil {
		if *value == 1 {
			return true
		}
	}
	return false
}
func NowAddDay() string {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = NowYmd()
	date, _ = time.Parse(layoutFormat, value)

	return date.AddDate(0, 0, 1).Format(layoutFormat)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}

func GetLastofCharacter(countChar int, char string) string {
	char = char[len(char)-countChar:]
	return char
}

func Encrypt(stringToEncrypt string) (encryptedString string) {
	res := b64.StdEncoding.EncodeToString([]byte(stringToEncrypt))
	return res
}

func Decrypt(encryptedString string, keyString string) (decryptedString string, err error) {
	res, err := b64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", models2.ErrGeneralMessage
	}
	return string(res), nil
}

func MetersToMils(meters float64) float64 {
	return meters * 0.000621371192
}
func MilsToMeters(mils float64) float64 {
	return mils * 1609.344
}

func MetersToKilometers(meters float64) float64 {
	return meters / 1000
}
func ValidationQueryParamsValueAble(request []models2.ValueAbleValidation) (bool, string) {

	for _, v := range request {
		found := InArray(v.Value,v.AvailableValue)
		if found == false{
			return false, v.Key + " " + models2.ErrInvalidValue.Error()
		}
	}
	return true, ""
}
func ValidationQueryParamsDataTypeNumberFloat(request []models2.DataTypeNumberFloatValidation) (bool, string) {

	for _, v := range request {
		if _, err := strconv.ParseFloat(v.Value, 64); err != nil {
			return false, v.Key + " " + models2.ErrInvalidDataType.Error()
		}
	}
	return true, ""
}
func ValidationQueryParamsDataTypeNumberInt(request []models2.DataTypeNumberIntValidation) (bool, string) {

	for _, v := range request {
		if _, err := strconv.Atoi(v.Value); err != nil {
			return false, v.Key + " " + models2.ErrInvalidDataType.Error()
		}
	}
	return true, ""
}

func ValidationQueryParamsDataTypeDate(request []models2.DataTypeNumberDateValidation) (bool, string) {
	for _, v := range request {
		if reflect.TypeOf(v.Value).String() == "string" && v.Value != "" {
			valueString := reflect.ValueOf(v.Value).String()
			convertDate := StringToDate(valueString)
			if convertDate.IsZero() == true {
				return false, v.Key + " " + models2.ErrInvalidDataType.Error()
			}
		}
	}
	return true, ""
}

func ValidationQueryParamsDataTypeDateMonth(request []models2.DataTypeNumberDateMonthValidation) (bool, string) {
	for _, v := range request {
		if reflect.TypeOf(v.Value).String() == "string" && v.Value != "" {
			valueString := reflect.ValueOf(v.Value).String()
			convertDate := StringToDateWithFormat(valueString,"2006-01")
			if convertDate.IsZero() == true {
				return false, v.Key + " " + models2.ErrInvalidDataType.Error()
			}
		}
	}
	return true, ""
}

func ValidationQueryParamsRequired(request []models2.RequiredValidation) (bool, string) {

	for _, v := range request {
		if v.Value == "" {
			return false, v.Key + " " + models2.ErrIsRequired.Error()
		}
	}
	return true, ""
}

func ValidationQueryParamsMaxMinLonglat(request []models2.MaxMinLonglatValidation) (bool, string) {

	for _, v := range request {
		if v.Key == "latitude_now" {
			lat := StringToFloat(v.Value)
			if lat < -90 {
				return false, v.Key + " " + models2.ErrInvalidValue.Error()
			}

			if lat > 90 {
				return false, v.Key + " " + models2.ErrInvalidValue.Error()
			}
		}

		if v.Key == "longitude_now" {
			lat := StringToFloat(v.Value)
			if lat < -180 {
				return false, v.Key + " " + models2.ErrInvalidValue.Error()
			}

			if lat > 180 {
				return false, v.Key + " " + models2.ErrInvalidValue.Error()
			}
		}
	}

	return true, ""
}
func ValidationQueryMaxMinNumber(validation []models2.MaxMinNumberValidation) (bool, string) {

	for _, v := range validation {
		val := StringToFloat(v.Value)
		if v.ValueMinNumber == -1 {
			continue
		}
		if val < v.ValueMinNumber {
			return false, v.Key + " " + models2.ErrInvalidValue.Error() + " min " + FloatToString(v.ValueMinNumber)
		}

		if val > v.ValueMaxNumber {
			return false, v.Key + " " + models2.ErrInvalidValue.Error() + " max " + FloatToString(v.ValueMaxNumber)
		}
	}

	return true, ""
}

func ValidationQueryParamsPageLimit(request []models2.PageLimitValidation) (bool, string) {

	for _, v := range request {
		if v.Key == "page" {
			p := StringToInt(v.Value)
			if p < 1 {
				return false, v.Key + " " + models2.ErrInvalidValue.Error()
			}
		}

		if v.Key == "limit" {
			l := StringToInt(v.Value)
			if l < 1 {
				return false, v.Key + " " + models2.ErrInvalidValue.Error()
			}

			if l > 100 {
				return false, v.Key + " " + models2.ErrInvalidValue.Error()
			}
		}
	}

	return true, ""
}

func GlobalValidationQueryParams(request models2.GlobalValidation) (bool, string) {
	checkQueryparams, message := ValidationQueryParamsRequired(request.RequiredValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsValueAble(request.ValueAbleValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsDataTypeDate(request.DataTypeNumberDateValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsDataTypeDateMonth(request.DataTypeNumberDateMonthValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsDataTypeNumberInt(request.DataTypeNumberIntValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsDataTypeNumberFloat(request.DataTypeNumberFloatValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsMaxMinLonglat(request.MaxMinLonglatValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsPageLimit(request.PageLimitValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryMaxMinNumber(request.MaxMinNumberValidation)
	if checkQueryparams == false {
		return false, message
	}

	return true, ""
}

func BindBody(reqbody io.ReadCloser) ([]byte, error) {
	var body []byte
	reqB, err := ioutil.ReadAll(reqbody)
	if err != nil {
		return nil, err
	}
	body = reqB

	return body, err
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

func JsonDecode(c echo.Context, request interface{}) (interface{}, error) {
	dec := json.NewDecoder(c.Request().Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func JsonString(object interface{}) string {
	res ,_:= json.Marshal(object)

	return string(res)
}
