package sample_module

import (
	"golang.org/x/net/context"
	"clean-architecture-beego/models"
)

type Usecase interface {
	GetList(ctx context.Context, page,limit,offset int) (res *models.SampleModulePaginationDto,message string,err error)
}
