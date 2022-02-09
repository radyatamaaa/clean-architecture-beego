package http_test

import (
	customerHttpHandler "clean-architecture-beego/internal/customer/delivery/http"
	_customerUsecaseMock "clean-architecture-beego/internal/customer/mocks"
	"clean-architecture-beego/internal/domain"
	"clean-architecture-beego/pkg/helpers/converter_value"
	testHelper "clean-architecture-beego/pkg/helpers/test"
	"context"
	"encoding/json"
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
	GroupUrl           = "/api/v1"
	GetCustomersUrl    = GroupUrl + "/customers"
	StoreCustomerUrl   = GroupUrl + "/customers"
	UpdateCustomerUrl  = GroupUrl + "/customers"
	DeleteCustomerUrl  = GroupUrl + "/customers"
	GetCustomerByIDUrl = GroupUrl + "/customers"
)

func TestCustomerHandler_LoadHandler(t *testing.T) {
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

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockCustomerUsecase := new(_customerUsecaseMock.Usecase)

		mockCustomerUsecase.On("GetProducts", mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).
			Return(mockCustomer, context.DeadlineExceeded)

		queryparam := "?offset=" + strconv.Itoa(offset) +
			"&limit=" + strconv.Itoa(limit)
		url := GetCustomersUrl + queryparam
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCustomerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.GetCustomers()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockCustomerUsecase.AssertExpectations(t)

	})

}

func TestCustomerHandler_StoreCustomer(t *testing.T) {
	//requestMock
	request := domain.CustomerStoreRequest{}
	err := faker.FakeData(&request)
	assert.NoError(t, err)

	//resultMock

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockCustomerUsecase := new(_customerUsecaseMock.Usecase)

		mockCustomerUsecase.On("SaveCustomer", mock.Anything,
			mock.AnythingOfType("domain.CustomerStoreRequest")).
			Return(nil)

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := StoreCustomerUrl
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCustomerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.StoreCustomer()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockCustomerUsecase.AssertExpectations(t)

	})

	t.Run("error-SaveCustomer-function", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("SaveCustomer", mock.Anything,
			mock.AnythingOfType("domain.CustomerStoreRequest")).
			Return(errors.New("Internel Server Error"))

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := StoreCustomerUrl
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.StoreCustomer()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("SaveCustomer", mock.Anything,
			mock.AnythingOfType("domain.CustomerStoreRequest")).
			Return(context.DeadlineExceeded)

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := StoreCustomerUrl
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.StoreCustomer()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-validate", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		request.CustomerName = "" //harusnya nil
		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := StoreCustomerUrl
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.StoreCustomer()

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

	})
}

func TestCustomerHandler_UpdateCustomer(t *testing.T) {
	//requestMock
	request := domain.CustomerUpdateRequest{}
	err := faker.FakeData(&request)
	assert.NoError(t, err)

	//resultMock

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockCustomerUsecase := new(_customerUsecaseMock.Usecase)

		mockCustomerUsecase.On("UpdateCustomer", mock.Anything,
			mock.AnythingOfType("domain.CustomerUpdateRequest")).
			Return(nil)

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := UpdateCustomerUrl
		req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCustomerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.UpdateCustomer()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockCustomerUsecase.AssertExpectations(t)

	})

	t.Run("error-UpdateCustomer-function", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("UpdateCustomer", mock.Anything,
			mock.AnythingOfType("domain.CustomerUpdateRequest")).
			Return(errors.New("Internel Server Error"))

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := UpdateCustomerUrl
		req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.UpdateCustomer()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("UpdateCustomer", mock.Anything,
			mock.AnythingOfType("domain.CustomerUpdateRequest")).
			Return(context.DeadlineExceeded)

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := UpdateCustomerUrl
		req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.UpdateCustomer()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	//t.Run("error-invalid-request-binding-object", func(t *testing.T) {
	//	//usecaseMock
	//	mockproductUsecase := new(_productUsecaseMock.Usecase)
	//
	//	requestBody := map[interface{}]interface{}{
	//		"name" : "test",
	//		"price" : "1",
	//		"stok" : "1",
	//	}
	//	j, err := json.Marshal(requestBody)
	//	assert.NoError(t, err)
	//	url := StoreProductUrl
	//	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
	//	assert.NoError(t, err)
	//	req.WithContext(context.Background())
	//
	//	rec := httptest.NewRecorder()
	//
	//	handler := productHttpHandler.ProductHandler{
	//		ProductUseCase: mockproductUsecase,
	//	}
	//	testHelper.PrepareHandler(t,&handler.Controller,req,rec)
	//	handler.Prepare()
	//
	//	handler.StoreProduct()
	//
	//	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	//
	//})

	t.Run("error-validate", func(t *testing.T) {
		//usecaseMock
		mockCustomerUsecase := new(_customerUsecaseMock.Usecase)

		request.Id = 0
		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := UpdateCustomerUrl
		req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCustomerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Prepare()

		handler.UpdateCustomer()

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

	})
}

func TestProductHandler_DeleteCustomer(t *testing.T) {
	//requestMock
	id := converter_value.IntToString(1)

	//resultMock

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("DeleteCustomer", mock.Anything,
			mock.AnythingOfType("int")).
			Return(nil)

		url := DeleteCustomerUrl + "/" + id
		req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Ctx.Input.SetParam("id", id)
		handler.Prepare()

		handler.DeleteCustomer()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-DeleteCustomer-function", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("DeleteCustomer", mock.Anything,
			mock.AnythingOfType("int")).
			Return(errors.New("Internel Server Error"))

		url := DeleteCustomerUrl + "/" + id
		req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Ctx.Input.SetParam("id", id)
		handler.Prepare()

		handler.DeleteCustomer()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("DeleteCustomer", mock.Anything,
			mock.AnythingOfType("int")).
			Return(context.DeadlineExceeded)

		url := DeleteCustomerUrl + "/" + id
		req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Ctx.Input.SetParam("id", id)
		handler.Prepare()

		handler.DeleteCustomer()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

}
func TestCustomerHandler_GetCustomerById(t *testing.T) {

	//resultMock
	mockCustomer := &domain.CustomerObjectResponse{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	//requestMock
	id := converter_value.IntToString(int(mockCustomer.Id))

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("GetCustomerById", mock.Anything,
			mock.AnythingOfType("uint")).
			Return(mockCustomer, nil)

		param := "/" + id
		url := GetCustomersUrl + param
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Ctx.Input.SetParam("id", id)
		handler.Prepare()

		handler.GetCustomerByID()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-GetCustomerById-function", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("GetCustomerById", mock.Anything,
			mock.AnythingOfType("uint")).
			Return(mockCustomer, errors.New("Internal Server Error"))

		url := GetCustomerByIDUrl + "/" + id
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Ctx.Input.SetParam("id", id)
		handler.Prepare()

		handler.GetCustomerByID()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockCostumerUsecase := new(_customerUsecaseMock.Usecase)

		mockCostumerUsecase.On("GetCustomerById", mock.Anything,
			mock.AnythingOfType("uint")).
			Return(mockCustomer, context.DeadlineExceeded)

		url := GetCustomerByIDUrl + "/" + id
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := customerHttpHandler.CustomerHandler{
			CustomerUseCase: mockCostumerUsecase,
		}
		testHelper.PrepareHandler(t, &handler.Controller, req, rec)
		handler.Ctx.Input.SetParam("id", id)
		handler.Prepare()

		handler.GetCustomerByID()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockCostumerUsecase.AssertExpectations(t)

	})
}
