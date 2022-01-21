package sample_module

import (
	"golang.org/x/net/context"
	"clean-architecture-beego/models"
)

type Repository interface {
	List(ctx context.Context, limit, offset int) (res []models.SampleModule, err error)
	Count(ctx context.Context) (res int, err error)
}
