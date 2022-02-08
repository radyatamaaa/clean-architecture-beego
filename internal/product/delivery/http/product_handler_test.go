package http_test

import (
	"clean-architecture-beego/internal/domain"
	productHttpHandler "clean-architecture-beego/internal/product/delivery/http"
	_productUsecaseMock "clean-architecture-beego/internal/product/mocks"
	"context"
	"encoding/json"
	"errors"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	testHelper "clean-architecture-beego/pkg/helpers/test"
)

var (
	GroupUrl = "/api/v1"
	GetProductsUrl           = GroupUrl + "/products"
	StoreProductUrl           = GroupUrl + "/products"
)

func TestProductHandler_LoadHandler(t *testing.T) {
	mockproductUsecase := new(_productUsecaseMock.Usecase)

	productHttpHandler.NewProductHandler(mockproductUsecase)

	mockproductUsecase.AssertExpectations(t)
}

func TestProductHandler_GetProducts(t *testing.T) {
	//requestMock
	offset := 0
	limit := 10

	//resultMock
	mockProduct := []domain.ProductObjectResponse{}
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("GetProducts", mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).
			Return(mockProduct, nil)


		queryparam := "?offset=" + strconv.Itoa(offset) +
			"&limit=" + strconv.Itoa(limit)
		url := GetProductsUrl + queryparam
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.GetProducts()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-GetProducts-function", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("GetProducts", mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).
			Return(mockProduct, errors.New("Internal Server Error"))


		queryparam := "?offset=" + strconv.Itoa(offset) +
			"&limit=" + strconv.Itoa(limit)
		url := GetProductsUrl + queryparam
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.GetProducts()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("GetProducts", mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).
			Return(mockProduct, context.DeadlineExceeded)


		queryparam := "?offset=" + strconv.Itoa(offset) +
			"&limit=" + strconv.Itoa(limit)
		url := GetProductsUrl + queryparam
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.GetProducts()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

}

func TestProductHandler_StoreProduct(t *testing.T) {
	//requestMock
	request := domain.ProductStoreRequest{}
	err := faker.FakeData(&request)
	assert.NoError(t, err)

	//resultMock


	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("SaveProduct", mock.Anything,
			mock.AnythingOfType("domain.ProductStoreRequest")).
			Return( nil)

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := StoreProductUrl
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.StoreProduct()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-SaveProduct-function", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("SaveProduct", mock.Anything,
			mock.AnythingOfType("domain.ProductStoreRequest")).
			Return( errors.New("Internel Server Error"))

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := StoreProductUrl
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.StoreProduct()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("SaveProduct", mock.Anything,
			mock.AnythingOfType("domain.ProductStoreRequest")).
			Return( context.DeadlineExceeded)

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := StoreProductUrl
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.StoreProduct()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-validate", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		request.Price = nil
		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := StoreProductUrl
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.StoreProduct()

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

	})

}
