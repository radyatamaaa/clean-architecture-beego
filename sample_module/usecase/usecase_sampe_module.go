package usecase

import (
	"context"
	"clean-architecture-beego/helper/logger"
	models2 "clean-architecture-beego/helper/models"
	"clean-architecture-beego/models"
	"clean-architecture-beego/sample_module"
	"time"
)

type SampleModuleUsecase struct {
	sampleModuleRepo    sample_module.Repository
	contextTimeout time.Duration
	log            logger.Logger
}

// NewSampleModuleUsecase will create new an SampleModuleUsecase object representation of sample_module.Usecase interface
func NewSampleModuleUsecase(sampleModuleRepo sample_module.Repository, log logger.Logger, timeout time.Duration) sample_module.Usecase {
	return &SampleModuleUsecase{
		sampleModuleRepo:    sampleModuleRepo,
		contextTimeout: timeout,
		log:            log,
	}
}

func (c SampleModuleUsecase) GetList(ctx context.Context, page,limit,offset int) (res *models.SampleModulePaginationDto,message string,err error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	list, err := c.sampleModuleRepo.List(ctx, limit, offset)
	if err != nil {
		c.log.Error("sample_module.usecase.SampleModuleUsecase.GetList: %s", err.Error())
		return nil, models2.ErrGeneralMessage.Error(), err
	}

	data := make([]models.SampleModuleDto,len(list))

	for i,val := range list{
		data[i] = data[i].MappingToDto(val)
	}

	totalRecords,err := c.sampleModuleRepo.Count(ctx)
	if err != nil {
		c.log.Error("sample_module.usecase.SampleModuleUsecase.GetList: %s", err.Error())
		return nil, models2.ErrGeneralMessage.Error(), err
	}

	result := models.SampleModulePaginationDto{}

	result = result.MappingToPaginatorDto(data,page,limit,totalRecords)


	return &result, models2.GeneralSuccess,nil
}
