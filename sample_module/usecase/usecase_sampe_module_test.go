package usecase_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"clean-architecture-beego/helper/logger"
	models2 "clean-architecture-beego/helper/models"
	"clean-architecture-beego/models"
	"clean-architecture-beego/sample_module/usecase"

	//_churnUsecaseMock "clean-architecture-beego/churn/mocks"
	_sampeModuleRepositoryMock "clean-architecture-beego/sample_module/mocks"
	"testing"
	"time"
)

var (
	timeoutContext = time.Second * 30
	l              = logger.L
)

func TestSampleModuleUsecase_GetList(t *testing.T) {

	//resultMock
	mockSampleModuleList := []models.SampleModule{}
	mockSampleModule := models.SampleModule{}
	mockSampleModule = mockSampleModule.MappingExpampleData()
	mockSampleModuleList = append(mockSampleModuleList, mockSampleModule)

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockSampleModuleRepo := new(_sampeModuleRepositoryMock.Repository)

		mockSampleModuleRepo.On("List", mock.Anything,
			mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockSampleModuleList, nil).Once()

		mockSampleModuleRepo.On("Count", mock.Anything).Return(10, nil).Once()

		u := usecase.NewSampleModuleUsecase(mockSampleModuleRepo, l, timeoutContext)

		a, _, err := u.GetList(context.TODO(), 1, 5, 0)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockSampleModuleRepo.AssertExpectations(t)
	})

	t.Run("error-List-function", func(t *testing.T) {
		//repositoryMock
		mockSampleModuleRepo := new(_sampeModuleRepositoryMock.Repository)

		mockSampleModuleRepo.On("List", mock.Anything,
			mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, models2.ErrInternalServerError).Once()

		u := usecase.NewSampleModuleUsecase(mockSampleModuleRepo, l, timeoutContext)

		a, _, err := u.GetList(context.TODO(), 1, 5, 0)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockSampleModuleRepo.AssertExpectations(t)
	})

	t.Run("error-Count-function", func(t *testing.T) {
		//repositoryMock
		mockSampleModuleRepo := new(_sampeModuleRepositoryMock.Repository)

		mockSampleModuleRepo.On("List", mock.Anything,
			mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockSampleModuleList, nil).Once()

		mockSampleModuleRepo.On("Count", mock.Anything).Return(0, models2.ErrInternalServerError).Once()

		u := usecase.NewSampleModuleUsecase(mockSampleModuleRepo, l, timeoutContext)

		a, _, err := u.GetList(context.TODO(), 1, 5, 0)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockSampleModuleRepo.AssertExpectations(t)
	})

}
