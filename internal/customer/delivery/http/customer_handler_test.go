package http_test

import (
	customerHttpHandler "clean-architecture-beego/internal/customer/delivery/http"
	_customerUsecaseMock "clean-architecture-beego/internal/customer/mocks"
	"clean-architecture-beego/internal/domain"
	testHelper "clean-architecture-beego/pkg/helpers/test"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	GroupUrl        = "/api/v1"
	GetCustomersUrl = GroupUrl + "/customers"
)

func TestProductHandler_LoadHandler(t *testing.T) {
	mockCustomerUsecase := new(_customerUsecaseMock.Usecase)

	customerHttpHandler.NewCustomerHandler(mockCustomerUsecase)

	mockCustomerUsecase.AssertExpectations(t)
}

func TestCustomerHandler_GetCustomers(t *testing.T) {
	//requestMock
	offset := 0
	limit := 10

	//resultMock
	mockCustomer := []domain.Customer{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("GetCustomers", mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).
			Return(mockCustomer, nil)

		queryparam := "?offset=" + strconv.Itoa(offset) +
			"&limit=" + strconv.Itoa(limit)
		url := GetCustomersUrl + queryparam
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.GetCustomers()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-GetCustomers-function", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("GetCustomers", mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).
			Return(mockCustomer, errors.New("Internal Server Error"))

		queryparam := "?offset=" + strconv.Itoa(offset) +
			"&limit=" + strconv.Itoa(limit)
		url := GetCustomersUrl + queryparam
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.GetCustomers()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

}

func TestCustomerHandler_GetCustomerById(t *testing.T) {
	//param
	idParam := 1

	//resultMock
	mockCustomer := &domain.Customer{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("GetCustomerById", mock.Anything,
			mock.AnythingOfType("uint")).
			Return(mockCustomer, nil)

		param := "/" + strconv.Itoa(idParam)
		url := GetCustomersUrl + param
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.GetCustomerByID()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-GetCustomerById-function", func(t *testing.T) {
		//param
		falseParam := "string"

		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("GetCustomerById", mock.Anything,
			mock.AnythingOfType("uint")).
			Return(mockCustomer, errors.New("Invalid Query"))

		param := "/" + falseParam
		url := GetCustomersUrl + param
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

	})
}
