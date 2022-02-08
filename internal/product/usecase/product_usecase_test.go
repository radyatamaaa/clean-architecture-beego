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
func TestProductUseCase_GetProducts(t *testing.T) {
	//resultMock
	mockProduct := []domain.Product{}
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	limit := 5
	offset := 0

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Fetch", mock.Anything,
			mock.AnythingOfType("int"),mock.AnythingOfType("int")).Return(mockProduct, nil).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository)

		a, err := u.GetProducts(context.TODO(), limit,offset)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-Fetch-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Fetch", mock.Anything,
			mock.AnythingOfType("int"),mock.AnythingOfType("int")).Return([]domain.Product{}, errors.New("invalid query")).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository)

		_, err := u.GetProducts(context.TODO(), limit,offset)

		assert.Error(t, err)

		mockProductRepository.AssertExpectations(t)
	})
}

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

	t.Run("error-FindByID-function", func(t *testing.T) {
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

func TestProductUseCase_SaveProduct(t *testing.T) {
	//requestMock
	mockProduct := domain.ProductStoreRequest{}
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Store", mock.Anything,
			mock.AnythingOfType("domain.Product")).Return(nil).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository)

		err := u.SaveProduct(context.TODO(), mockProduct)

		assert.NoError(t, err)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-Store-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Store", mock.Anything,
			mock.AnythingOfType("domain.Product")).Return(errors.New("invalid query")).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository)

		err := u.SaveProduct(context.TODO(), mockProduct)

		assert.Error(t, err)

		mockProductRepository.AssertExpectations(t)
	})
}

func TestProductUseCase_UpdateProduct(t *testing.T) {
	//resultMock
	mockProduct := domain.ProductUpdateRequest{}
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)


	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Update", mock.Anything,
			mock.AnythingOfType("domain.Product")).Return(nil).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository)

		err := u.UpdateProduct(context.TODO(), mockProduct)

		assert.NoError(t, err)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-Store-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Update", mock.Anything,
			mock.AnythingOfType("domain.Product")).Return(errors.New("invalid query")).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository)

		err := u.UpdateProduct(context.TODO(), mockProduct)

		assert.Error(t, err)

		mockProductRepository.AssertExpectations(t)
	})
}
