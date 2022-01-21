package models

import "clean-architecture-beego/models"

type SwaggerListSuccess struct {
	StatusCode int                               `json:"status_code" example:"200"`
	Status     string                            `json:"status_desc" example:"OK"`
	Msg        string                            `json:"message" example:"Success"`
	Data       *models.SampleModulePaginationDto `json:"data"`
	Errors     *string                           `json:"errors" example:"null"`
}
