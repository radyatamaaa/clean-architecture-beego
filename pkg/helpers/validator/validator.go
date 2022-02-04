package validator

import (
	"clean-architecture-beego/pkg/helpers/converter_value"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/common/log"
	"reflect"
	"strings"
)
func NewValidator() *GlobalValidator {
	v :=  &GlobalValidator{
		Validator: validator.New(),
	}
	return v
}
func (v *GlobalValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

func (v *GlobalValidator) RegisterValidators() error {
	if err := v.Validator.RegisterValidation("validate_required_if_another_filed", validateRequireIfAnotherField); err != nil {
		log.Error(err)
	}

	return nil
}
func validateRequireIfAnotherField(fl validator.FieldLevel) bool {
	param := strings.Split(fl.Param(), `:`)
	paramField := param[0]
	paramValue := param[1]

	if paramField == `` {
		return true
	}

	// param field reflect.Value.
	var paramFieldValue reflect.Value

	if fl.Parent().Kind() == reflect.Ptr {
		paramFieldValue = fl.Parent().Elem().FieldByName(paramField)
	} else {
		paramFieldValue = fl.Parent().FieldByName(paramField)
	}

	if isEq(paramFieldValue, paramValue) == false {
		return true
	}
	return hasValue(fl)
}

func isEq(field reflect.Value, value string) bool {
	switch field.Kind() {

	case reflect.String:
		return field.String() == value

	case reflect.Slice, reflect.Map, reflect.Array:
		p := converter_value. AsInt(value)

		return int64(field.Len()) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := converter_value. AsInt(value)

		return field.Int() == p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := converter_value. AsUint(value)

		return field.Uint() == p

	case reflect.Float32, reflect.Float64:
		p := converter_value. AsFloat(value)

		return field.Float() == p
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func hasValue(fl validator.FieldLevel) bool {
	return requireCheckFieldKind(fl, "")
}
func requireCheckFieldKind(fl validator.FieldLevel, param string) bool {
	field := fl.Field()
	if len(param) > 0 {
		if fl.Parent().Kind() == reflect.Ptr {
			field = fl.Parent().Elem().FieldByName(param)
		} else {
			field = fl.Parent().FieldByName(param)
		}
	}
	switch field.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !field.IsNil()
	default:
		_, _, nullable := fl.ExtractType(field)
		if nullable && field.Interface() != nil {
			return true
		}
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func PanicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}

