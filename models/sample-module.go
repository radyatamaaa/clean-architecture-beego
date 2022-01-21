package models

import "clean-architecture-beego/helper/models"

type SampleModule struct {
	Id int `json:"id"`
	Test string `json:"test"`
}

type SampleModuleDto struct {
	Id int `json:"id"`
	Test string `json:"test"`
}

type SampleModulePaginationDto struct {
	Data      []SampleModuleDto `json:"data"`
	Paginator models.Paginator  `json:"paginator"`
}

func (sm SampleModule)MappingExpampleData() SampleModule {
	sm.Id = 1
	sm.Test = "1adadasd"
	return sm
}

func (smd SampleModuleDto)MappingToDto(sm SampleModule) SampleModuleDto {
	smd.Id = sm.Id
	smd.Test = sm.Test
	return smd
}

func (sm SampleModulePaginationDto)MappingToPaginatorDto(data []SampleModuleDto,page,limit,totalAllRecords int) SampleModulePaginationDto {
	sm.Data = data
	sm.Paginator = sm.Paginator.MappingPaginator(page,limit,totalAllRecords,len(data))
	return sm
}

func (sm SampleModulePaginationDto)MappingExpampleData() SampleModulePaginationDto {
	sm.Data = []SampleModuleDto{
		SampleModuleDto{Test: "asdasd",Id: 1},
	}
	sm.Paginator = sm.Paginator.MappingPaginator(1,1,1,1)
	return sm
}




