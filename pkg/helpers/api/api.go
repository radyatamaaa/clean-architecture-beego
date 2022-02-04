package api

import (
	"errors"
	"net/http"
)
var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrUnAuthorize = errors.New("Unauthorize")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist or duplicate")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Bad Request")

	ErrForbidden = errors.New("you don't have permission to access this resource")
)
func (r *Response) MappingResponseSuccess(message string, data interface{}) {
	r.StatusCode = http.StatusOK
	r.Status = http.StatusText(r.StatusCode)
	r.Msg = message
	r.Data = data
	r.Errors = nil
}

func (r *Response) MappingResponseError(statusCode int, message string, error interface{}) {
	r.StatusCode = statusCode
	r.Status = http.StatusText(r.StatusCode)
	r.Msg = message
	r.Data = nil
	r.Errors = error
}

func (r *Response) MappingResponseErrorValidation(statusCode int, message string, error interface{}) {
	r.StatusCode = statusCode
	r.Status = http.StatusText(r.StatusCode)
	r.Msg = message
	r.Data = nil
	r.Errors = error
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrForbidden:
		return http.StatusForbidden
	case ErrNotFound:
		return http.StatusNotFound
	case ErrUnAuthorize:
		return http.StatusUnauthorized
	case ErrConflict:
		return http.StatusConflict
	case ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

