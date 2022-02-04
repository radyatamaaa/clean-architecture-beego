package api

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status_desc"`
	Msg        string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"errors"`
	ErrorValidation []ErrorValidation `json:"error_validation"`
}

type ErrorValidation struct {
	Column string `json:"column"`
	Message string `json:"message"`
}
