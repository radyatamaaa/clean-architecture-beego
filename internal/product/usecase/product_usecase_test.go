package usecase_test

import (
	"clean-architecture-beego/internal/domain"
	_mockProductRepository "clean-architecture-beego/internal/product/mocks"
	"clean-architecture-beego/internal/product/usecase"
	"context"
	"errors"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)
var (
	timeoutContext = time.Second * 30
)
func TestProductUseCase_GetProductById(t *testing.T) {
	//resultMock
	mockProduct := domain.Product{}
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	id := mockProduct.Id

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("FindByID", mock.Anything,
			mock.AnythingOfType("uint")).Return(mockProduct, nil).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository)

		a, err := u.GetProductById(context.TODO(), id)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-GetByIdNumber-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("FindByID", mock.Anything,
			mock.AnythingOfType("uint")).Return(domain.Product{}, errors.New("invalid query")).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository)

		_, err := u.GetProductById(context.TODO(), id)

		assert.Error(t, err)

		mockProductRepository.AssertExpectations(t)
	})
}
