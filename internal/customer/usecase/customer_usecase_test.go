package usecase_test

import (
	_mockCustomerRepository "clean-architecture-beego/internal/customer/mocks"
	"clean-architecture-beego/internal/customer/usecase"
	"clean-architecture-beego/internal/domain"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	timeoutContext = time.Second * 30
)

func TestCustomerUseCase_GetCustomer(t *testing.T) {
	//resultMock
	mockCustomer := []domain.Customer{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	limit := 5
	offset := 0

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("Fetch", mock.Anything,
			mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockCustomer, nil).Once() //type int

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		a, err := u.GetCustomers(context.TODO(), limit, offset)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockCustomerRepository.AssertExpectations(t)
	})

	t.Run("error-Fetch-function", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("Fetch", mock.Anything,
			mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]domain.Customer{}, errors.New("invalid query")).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		_, err := u.GetCustomers(context.TODO(), limit, offset)

		assert.Error(t, err)

		mockCustomerRepository.AssertExpectations(t)
	})
}

func TestCustomerUseCase_GetCustomerById(t *testing.T) {
	//resultMock
	mockCustomer := domain.Customer{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	id := mockCustomer.Id

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("FindByID", mock.Anything,
			mock.AnythingOfType("uint")).Return(mockCustomer, nil).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		a, err := u.GetCustomerById(context.TODO(), id)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockCustomerRepository.AssertExpectations(t)
	})

	t.Run("error-FindByID-function", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("FindByID", mock.Anything,
			mock.AnythingOfType("uint")).Return(domain.Customer{}, errors.New("invalid query")).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		_, err := u.GetCustomerById(context.TODO(), id)

		assert.Error(t, err)

		mockCustomerRepository.AssertExpectations(t)
	})
}

func TestCustomerUseCase_SaveCustomer(t *testing.T) {
	//requestMock
	mockCustomer := domain.CustomerStoreRequest{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("Store", mock.Anything,
			mock.AnythingOfType("domain.Customer")).Return(nil).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		err := u.SaveCustomer(context.TODO(), mockCustomer)

		assert.NoError(t, err)

		mockCustomerRepository.AssertExpectations(t)
	})

	t.Run("error-Store-function", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("Store", mock.Anything,
			mock.AnythingOfType("domain.Customer")).Return(errors.New("invalid query")).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		err := u.SaveCustomer(context.TODO(), mockCustomer)

		assert.Error(t, err)

		mockCustomerRepository.AssertExpectations(t)
	})
}

func TestCustomerUseCase_UpdateCustomer(t *testing.T) {
	//resultMock
	mockCustomer := domain.CustomerUpdateRequest{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("Update", mock.Anything,
			mock.AnythingOfType("domain.Customer")).Return(nil).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		err := u.UpdateCustomer(context.TODO(), mockCustomer)

		assert.NoError(t, err)

		mockCustomerRepository.AssertExpectations(t)
	})

	t.Run("error-Store-function", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("Update", mock.Anything,
			mock.AnythingOfType("domain.Customer")).Return(errors.New("invalid query")).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		err := u.UpdateCustomer(context.TODO(), mockCustomer)

		assert.Error(t, err)

		mockCustomerRepository.AssertExpectations(t)
	})
}

func TestCustomerUseCase_DeleteCustomer(t *testing.T) {
	//resultMock
	mockCustomer := domain.Customer{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	id := mockCustomer.Id

	t.Run("success", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("Delete", mock.Anything,
			mock.AnythingOfType("uint")).Return(mockCustomer, nil).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		err := u.DeleteCustomer(context.TODO(), id)

		assert.NoError(t, err)

		mockCustomerRepository.AssertExpectations(t)
	})

	t.Run("error-Delete-function", func(t *testing.T) {
		//repositoryMock
		mockCustomerRepository := new(_mockCustomerRepository.Repository)

		mockCustomerRepository.On("Delete", mock.Anything,
			mock.AnythingOfType("uint")).Return(domain.Customer{}, errors.New("invalid query")).Once()

		u := usecase.NewCustomerUseCase(timeoutContext, mockCustomerRepository)

		err := u.DeleteCustomer(context.TODO(), id)

		assert.Error(t, err)

		mockCustomerRepository.AssertExpectations(t)
	})
}
