package usecase_test

import (
	"clean-architecture-beego/internal/domain"
	_mockProductRepository "clean-architecture-beego/internal/product/mocks"
	"clean-architecture-beego/internal/product/usecase"
	"clean-architecture-beego/pkg/logger"
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
	l = logger.NewStdOutLogger(30,"all","Local",true,"1.1.0","kreditmu","clean-architecture-beego",logger.XmodeTest)
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
			mock.AnythingOfType("int"),mock.AnythingOfType("int")).Return(mockProduct, nil,"").Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		a, err ,_:= u.GetProducts(context.TODO(), limit,offset)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-Fetch-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Fetch", mock.Anything,
			mock.AnythingOfType("int"),mock.AnythingOfType("int")).Return([]domain.Product{}, errors.New("invalid query"),errors.New("invalid query").Error()).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		_, err ,_:= u.GetProducts(context.TODO(), limit,offset)

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
			mock.AnythingOfType("uint")).Return(mockProduct, nil,"").Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		a, err ,_:= u.GetProductById(context.TODO(), id)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-FindByID-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("FindByID", mock.Anything,
			mock.AnythingOfType("uint")).Return(domain.Product{}, errors.New("invalid query"),errors.New("invalid query").Error()).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		_, err,_ := u.GetProductById(context.TODO(), id)

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
			mock.AnythingOfType("domain.Product")).Return(nil,"").Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		err,_ := u.SaveProduct(context.TODO(), mockProduct)

		assert.NoError(t, err)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-Store-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Store", mock.Anything,
			mock.AnythingOfType("domain.Product")).Return(errors.New("invalid query"),errors.New("invalid query").Error()).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		err ,_:= u.SaveProduct(context.TODO(), mockProduct)

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
			mock.AnythingOfType("domain.Product")).Return(nil,"").Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		err,_ := u.UpdateProduct(context.TODO(), mockProduct)

		assert.NoError(t, err)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-Store-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Update", mock.Anything,
			mock.AnythingOfType("domain.Product")).Return(errors.New("invalid query"),errors.New("invalid query").Error()).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		err,_ := u.UpdateProduct(context.TODO(), mockProduct)

		assert.Error(t, err)

		mockProductRepository.AssertExpectations(t)
	})
}

func TestProductUseCase_DeleteProduct(t *testing.T) {
	//resultMock
	mockProduct := domain.Product{}
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	id := mockProduct.Id


	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Delete", mock.Anything,
			mock.AnythingOfType("int")).Return(nil,"").Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		err ,_:= u.DeleteProduct(context.TODO(), int(id))

		assert.NoError(t, err)

		mockProductRepository.AssertExpectations(t)
	})

	t.Run("error-Store-function", func(t *testing.T) {
		//repositoryMock
		mockProductRepository := new(_mockProductRepository.Repository)

		mockProductRepository.On("Delete", mock.Anything,
			mock.AnythingOfType("int")).Return(errors.New("invalid query"),errors.New("invalid query").Error()).Once()

		u := usecase.NewProductUseCase(timeoutContext,mockProductRepository,l)

		err ,_:= u.DeleteProduct(context.TODO(), int(id))

		assert.Error(t, err)

		mockProductRepository.AssertExpectations(t)
	})
}
