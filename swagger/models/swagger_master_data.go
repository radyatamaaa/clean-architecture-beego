package models

type SwaggerListStatusKetertarikanSuccess struct {
	StatusCode int         `json:"status_code" example:"200"`
	Status     string      `json:"status_desc" example:"OK"`
	Msg        string      `json:"message" example:"Success"`
	Data       []*struct{
		Id 		int `json:"id" example:"1"`
		Description string `json:"description" example:"Agen Tertarik"`
	} `json:"data"`
	Errors     *string `json:"errors" example:"null"`
}

