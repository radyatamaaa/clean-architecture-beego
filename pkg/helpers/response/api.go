package response

import (
	"encoding/json"
	"errors"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

const (
	InvalidApiKeyError    = "INVALID_API_KEY"
	UnauthorizedError     = "UNAUTHORIZED"
	RequestForbiddenError = "REQUEST_FORBIDDEN_ERROR"
	ApiValidationError    = "API_VALIDATION_ERROR"
	ResourceNotFoundError = "RESOURCE_NOT_FOUND"
	ServerError           = "SERVER_ERROR"
	RequestTimeout        = "REQUEST_TIMEOUT"
	InvalidCredentials    = "INVALID_CREDENTIAL"
)

var (
	// ErrApiValidationError will throw if request is invalid
	ErrApiValidationError = errors.New("invalid request, errors arise when your request has invalid parameters")
	// ErrUnAuthorize will throw if user is not authorize
	ErrUnAuthorize = errors.New("unauthorized")
	// ErrInvalidApiKeyError will throw if user no valid api key provided
	ErrInvalidApiKeyError = errors.New("no valid api key provided")
	// ErrRequestForbiddenError will throw if user don't have permission access the server
	ErrRequestForbiddenError = errors.New("you don't have permission to access this resource")
	// ErrResourceNotFoundError will throw if resource on the server not found
	ErrResourceNotFoundError = errors.New("the requested resources doesn't exist")
	// ErrRequestTimeout will throw if any the Internal Server Error happen
	ErrRequestTimeout = errors.New("the request to server is timeout, please try again")
	// ErrServerError will throw if any the Internal Server Error happen
	ErrServerError = errors.New("internal server error")
	// ErrInvalidCredential will throw if any the Internal Server Error happen
	ErrInvalidCredential = errors.New("invalid credential, please check your email or username or password")
)

type ApiResponse struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []Errors    `json:"errors,omitempty"`
}

type Errors struct {
	Field       string `json:"field"`
	Description string `json:"message"`
}

func (r *ApiResponse) Ok(ctx *context.Context, data interface{}) error {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(http.StatusOK)

	return ctx.Output.JSON(ApiResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}, beego.BConfig.RunMode != "prod", false)
}

func (r *ApiResponse) ErrorResponse(ctx *context.Context, httpStatus int, errorCode string, err error) error {
	var apiResponse ApiResponse
	var errorValidations []Errors

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(httpStatus)

	if ctx.Input.RequestBody != nil {
		validateJsonError := r.checkJsonRequest(err)
		if len(validateJsonError) != 0 {
			errorValidations = validateJsonError
		} else {
			if fields, ok := err.(validator.ValidationErrors); ok {
				for _, v := range fields {
					errorValidations = append(errorValidations, Errors{
						Field:       v.Field(),
						Description: fmt.Sprintf("Parameter %s %s", v.Field(), v.Error()),
					})
				}
			}
		}
	}

	apiResponse.Code = errorCode
	apiResponse.Message = r.getMessageErrorCode(errorCode)
	if len(errorValidations) > 0 {
		apiResponse.Errors = errorValidations
	}

	return ctx.Output.JSON(apiResponse, beego.BConfig.RunMode != "prod", false)
}

// checkJsonRequest Response API
func (r *ApiResponse) checkJsonRequest(err error) (response []Errors) {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var invalidUnmarshalError *json.InvalidUnmarshalError
	switch {
	case errors.As(err, &syntaxError):
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		errorValidation := Errors{
			Field:       "json",
			Description: msg,
		}
		response = append(response, errorValidation)
		return
	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := fmt.Sprintf("Request body contains badly-formed JSON")
		errorValidation := Errors{
			Field:       "json",
			Description: msg,
		}
		response = append(response, errorValidation)
		return
	case errors.As(err, &unmarshalTypeError):
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			errorValidation := Errors{
				Field:       ute.Field,
				Description: fmt.Sprintf("Parameter %s is invalid (type: %s)", ute.Field, ute.Type),
			}
			response = append(response, errorValidation)
			return
		}
	case errors.As(err, &invalidUnmarshalError):
		if ute, ok := err.(*json.InvalidUnmarshalError); ok {
			errorValidation := Errors{
				Field:       ute.Type.Name(),
				Description: ute.Error(),
			}
			response = append(response, errorValidation)
			return
		}
	}
	return response
}

func (r *ApiResponse) getErrorCode(err error) string {
	switch err {
	case ErrInvalidApiKeyError:
		return InvalidApiKeyError
	case ErrUnAuthorize:
		return UnauthorizedError
	case ErrRequestForbiddenError:
		return RequestForbiddenError
	case ErrApiValidationError:
		return ApiValidationError
	case ErrResourceNotFoundError:
		return ResourceNotFoundError
	default:
		return ServerError
	}
}

func (r *ApiResponse) getMessageErrorCode(errorCode string) string {
	switch errorCode {
	case InvalidApiKeyError:
		return ErrInvalidApiKeyError.Error()
	case UnauthorizedError:
		return ErrUnAuthorize.Error()
	case RequestForbiddenError:
		return ErrRequestForbiddenError.Error()
	case ApiValidationError:
		return ErrApiValidationError.Error()
	case ResourceNotFoundError:
		return ErrResourceNotFoundError.Error()
	case RequestTimeout:
		return ErrRequestTimeout.Error()
	case InvalidCredentials:
		return ErrInvalidCredential.Error()
	default:
		return ErrServerError.Error()
	}
}
