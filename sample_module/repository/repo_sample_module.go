package repository

import (
	"context"
	"gorm.io/gorm"
	"clean-architecture-beego/helper"
	"clean-architecture-beego/helper/logger"
	"clean-architecture-beego/models"

	"clean-architecture-beego/sample_module"
)

type SampleModuleRepository struct {
	Conn *gorm.DB
	log  logger.Logger
}

// NewSampleModuleRepository will create an object that represent the sample_module.Repository interface
func NewSampleModuleRepository(Conn *gorm.DB, log logger.Logger) sample_module.Repository {
	return &SampleModuleRepository{Conn, log}
}

//query
//func (n SampleModuleRepository) query(ctx context.Context) {
//	panic("implement me")
//}
//
////messaging
//func (n SampleModuleRepository) messaging(ctx context.Context) {
//	panic("implement me")
//}
//
////grpc
//func (n SampleModuleRepository) grpc(ctx context.Context) {
//	panic("implement me")
//}

//functional
func (n SampleModuleRepository) List(ctx context.Context, limit, offset int) (res []models.SampleModule, err error) {
	res = make([]models.SampleModule, 10)
	for i, _ := range res {
		res[i].Id = i + 1
		res[i].Test = helper.RandomString(10)
	}
	if offset > len(res) {
		offset = len(res)
	}

	end := offset + limit
	if end > len(res) {
		end = len(res)
	}

	return res[offset:end], nil
}

func (n SampleModuleRepository) Count(ctx context.Context) (res int, err error) {
	return 10, nil
}
