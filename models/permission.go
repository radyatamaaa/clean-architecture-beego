package models

import (
	"github.com/labstack/echo/v4"
	"clean-architecture-beego/helper"
	"clean-architecture-beego/helper/models"
	"time"
)

type Permission struct {
	ID          string    `gorm:"type:varchar(60);column:id;primary_key:true"`
	Feature     string    `gorm:"type:varchar(50);column:feature"`
	Url         string    `gorm:"type:varchar(255);column:url"`
	Description string    `gorm:"type:varchar(100);column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type PermissionListRequest struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	Page int `json:"page"`
}

type PermissionListParam struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

type PermissionDto struct {
	ID          string `json:"id"`
	Feature     string `json:"feature"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type PermissionPaginationDto struct {
	Data      []*PermissionDto `json:"data"`
	Paginator models.Paginator  `json:"paginator"`
}


//mapping
func (r *PermissionListRequest)MappingFromContextRequest(c echo.Context) *PermissionListRequest {
	result := &PermissionListRequest{
		Limit:  helper.StringToInt(c.QueryParam("limit")),
		Offset: 0,
		Page:   helper.StringToInt(c.QueryParam("page")),
	}
	return result
}
func (r *PermissionListRequest)MappingToPermissionListParam() *PermissionListParam {
	result := &PermissionListParam{
		Limit:  r.Limit,
		Offset: r.Offset,
	}
	return result
}

func (smd *PermissionDto)MappingToDto(sm *Permission) *PermissionDto {
	smd = &PermissionDto{
		ID:          sm.ID,
		Feature:     sm.Feature,
		Url:         sm.Url,
		Description: sm.Description,
	}
	return smd
}

func (sm *PermissionPaginationDto)MappingToPaginatorDto(data []*PermissionDto,page,limit,totalAllRecords int) *PermissionPaginationDto {
	sm.Data = data
	sm.Paginator = sm.Paginator.MappingPaginator(page,limit,totalAllRecords,len(data))
	return sm
}

func (sm PermissionPaginationDto)MappingExpampleData() PermissionPaginationDto {
	sm.Data = []*PermissionDto{
		{
			ID:          "01346bcb-94e4-4aa9-98a5-bf375d730bfe",
			Feature:     "Customer Unlock List",
			Url:         "/clean-architecture-beego/admin/contact-center/unlock/list",
			Description: "List of customers unlock",
		},
	}
	sm.Paginator = sm.Paginator.MappingPaginator(1,1,1,1)
	return sm
}

func (sm *Permission)MappingExpampleData() *Permission {
	result := &Permission{
		ID:          "01346bcb-94e4-4aa9-98a5-bf375d730bfe",
		Feature:     "Customer Unlock List",
		Url:         "/clean-architecture-beego/admin/contact-center/unlock/list",
		Description: "List of customers unlock",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return result
}

func (sm *PermissionListParam)MappingExpampleData() *PermissionListParam {
	result := &PermissionListParam{
		Limit:  1,
		Offset: 0,
	}
	return result
}

func (sm *PermissionListRequest)MappingExpampleData() *PermissionListRequest {
	result := &PermissionListRequest{
		Limit:  1,
		Offset: 0,
		Page: 1,
	}
	return result
}