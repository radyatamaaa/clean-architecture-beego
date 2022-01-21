package models

type SwaggerErrorNotFound struct {
	StatusCode int       `json:"status_code" example:"404"`
	Status     string    `json:"status_desc" example:"Not Found"`
	Msg        string    `json:"message" example:"<Error-Message-For-Client-User>"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"<Error-Message-For-System>"`
}

type SwaggerErrorConflict struct {
	StatusCode int       `json:"status_code" example:"409"`
	Status     string    `json:"status_desc" example:"Conflict"`
	Msg        string    `json:"message" example:"<Error-Message-For-Client-User>"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"<Error-Message-For-System>"`
}

type SwaggerErrorUnAuthorize struct {
	StatusCode int       `json:"status_code" example:"401"`
	Status     string    `json:"status_desc" example:"Unauthorized"`
	Msg        string    `json:"message" example:"<Error-Message-For-Client-User>"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"<Error-Message-For-System>"`
}

type SwaggerErrorBadParamInput struct {
	StatusCode int       `json:"status_code" example:"400"`
	Status     string    `json:"status_desc" example:"Bad Request"`
	Msg        string    `json:"message" example:"<Error-Message-For-Client-User>"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"<Error-Message-For-System>"`
}

type SwaggerOtherInternalServerError struct {
	StatusCode int       `json:"status_code" example:"500"`
	Status     string    `json:"status_desc" example:"Internal Server Error"`
	Msg        string    `json:"message" example:"<Error-Message-For-Client-User>"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"<Error-Message-For-System>"`
}

type SwaggerErrorInternalServerError struct {
	StatusCode int       `json:"status_code" example:"500"`
	Status     string    `json:"status_desc" example:"Internal Server Error"`
	Msg        string    `json:"message" example:"something wrong"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"the error that can be obtained from the (database server,cache server,other microservices shutdown) or any error code"`
}

type SwaggerErrorTokenUnAuthorize struct {
	StatusCode int       `json:"status_code" example:"401"`
	Status     string    `json:"status_desc" example:"Unauthorized"`
	Msg        string    `json:"message" example:"Invalid authorization token"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"null"`
}

type SwaggerErrorExpiredTokenUnAuthorize struct {
	StatusCode int       `json:"status_code" example:"401"`
	Status     string    `json:"status_desc" example:"Unauthorized"`
	Msg        string    `json:"message" example:"Authorization token has expired"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"null"`
}

type SwaggerErrorTokenForbiddenPermission struct {
	StatusCode int       `json:"status_code" example:"403"`
	Status     string    `json:"status_desc" example:"Forbidden"`
	Msg        string    `json:"message" example:"you don't have permission to access this resource"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"null"`
}

type SwaggerErrorExpiredTokenUnAuthorizeAndTokenUnAuthorize struct {
	ErrorTokenUnAuthorize        SwaggerErrorTokenUnAuthorize        `json:"error_token_un_authorize"`
	ErrorExpiredTokenUnAuthorize SwaggerErrorExpiredTokenUnAuthorize `json:"error_expired_token_un_authorize"`
}

type SwaggerErrorNotFoundRoutes struct {
	StatusCode int       `json:"status_code" example:"404"`
	Status     string    `json:"status_desc" example:"Not Found"`
	Msg        string    `json:"message" example:"Routes Not Found"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"null"`
}

type SwaggerErrorMethodNotAllowed struct {
	StatusCode int       `json:"status_code" example:"405"`
	Status     string    `json:"status_desc" example:"Method Not Allowed"`
	Msg        string    `json:"message" example:"Method Not Allowed"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"null"`
}

type SwaggerErrorInvalidMethod struct {
	StatusCode int       `json:"status_code" example:"400"`
	Status     string    `json:"status_desc" example:"Bad Request"`
	Msg        string    `json:"message" example:"Invalid Method"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"null"`
}
