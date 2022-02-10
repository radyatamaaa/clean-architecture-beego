package http_test

import (
	"clean-architecture-beego/internal/domain"
	productHttpHandler "clean-architecture-beego/internal/product/delivery/http"
	_productUsecaseMock "clean-architecture-beego/internal/product/mocks"
	"clean-architecture-beego/pkg/helpers/converter_value"
	testHelper "clean-architecture-beego/pkg/helpers/test"
	"clean-architecture-beego/pkg/logger"
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
)

var (
	GroupUrl = "/api/v1"
	GetProductsUrl           = GroupUrl + "/products"
	StoreProductUrl           = GroupUrl + "/products"
	UpdateProductUrl           = GroupUrl + "/products"
	DeleteProductUrl           = GroupUrl + "/products"
	GetProductByIDUrl           = GroupUrl + "/products"
)
var (
	l = logger.NewStdOutLogger(30,"all","Local",true)
)
func TestProductHandler_LoadHandler(t *testing.T) {
	mockproductUsecase := new(_productUsecaseMock.Usecase)

	productHttpHandler.NewProductHandler(mockproductUsecase,l)

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

func TestProductHandler_UpdateProduct(t *testing.T) {
	//requestMock
	request := domain.ProductUpdateRequest{}
	err := faker.FakeData(&request)
	assert.NoError(t, err)

	//resultMock


	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("UpdateProduct", mock.Anything,
			mock.AnythingOfType("domain.ProductUpdateRequest")).
			Return( nil)

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := UpdateProductUrl
		req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.UpdateProduct()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-UpdateProduct-function", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("UpdateProduct", mock.Anything,
			mock.AnythingOfType("domain.ProductUpdateRequest")).
			Return( errors.New("Internel Server Error"))

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := UpdateProductUrl
		req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.UpdateProduct()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("UpdateProduct", mock.Anything,
			mock.AnythingOfType("domain.ProductUpdateRequest")).
			Return( context.DeadlineExceeded)

		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := UpdateProductUrl
		req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.UpdateProduct()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockproductUsecase.AssertExpectations(t)

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
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		request.Id = 0
		j, err := json.Marshal(request)
		assert.NoError(t, err)
		url := UpdateProductUrl
		req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(j)))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.UpdateProduct()

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

	})
}

func TestProductHandler_DeleteProduct(t *testing.T) {
	//requestMock
	id := converter_value.IntToString(1)

	//resultMock


	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("DeleteProduct", mock.Anything,
			mock.AnythingOfType("int")).
			Return( nil)

		url := DeleteProductUrl + "/" +id
		req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Ctx.Input.SetParam("id",id)
		handler.Prepare()

		handler.DeleteProduct()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-DeleteProduct-function", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("DeleteProduct", mock.Anything,
			mock.AnythingOfType("int")).
			Return( errors.New("Internel Server Error"))


		url := DeleteProductUrl + "/" +id
		req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Ctx.Input.SetParam("id",id)
		handler.Prepare()

		handler.DeleteProduct()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("DeleteProduct", mock.Anything,
			mock.AnythingOfType("int")).
			Return( context.DeadlineExceeded)


		url := DeleteProductUrl + "/" +id
		req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Ctx.Input.SetParam("id",id)
		handler.Prepare()

		handler.DeleteProduct()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

}

func TestProductHandler_GetProductByID(t *testing.T) {

	//resultMock
	mockProduct := &domain.ProductObjectResponse{}
	err := faker.FakeData(mockProduct)
	assert.NoError(t, err)

	//requestMock
	id := converter_value.IntToString(int(mockProduct.Id))

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("GetProductById", mock.Anything,
			mock.AnythingOfType("uint")).
			Return(mockProduct, nil)

		url := GetProductByIDUrl + "/" +id
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Ctx.Input.SetParam("id",id)
		handler.Prepare()

		handler.GetProductByID()

		assert.Equal(t, http.StatusOK, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-GetProducts-function", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("GetProductById", mock.Anything,
			mock.AnythingOfType("uint")).
			Return(mockProduct, errors.New("Internal Server Error"))


		url := GetProductByIDUrl + "/" +id
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Ctx.Input.SetParam("id",id)
		handler.Prepare()

		handler.GetProductByID()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

	t.Run("error-timeout", func(t *testing.T) {
		//usecaseMock
		mockproductUsecase := new(_productUsecaseMock.Usecase)

		mockproductUsecase.On("GetProductById", mock.Anything,
			mock.AnythingOfType("uint")).
			Return(mockProduct, context.DeadlineExceeded)


		url := GetProductByIDUrl + "/" +id
		req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
		assert.NoError(t, err)
		req.WithContext(context.Background())

		rec := httptest.NewRecorder()

		handler := productHttpHandler.ProductHandler{
			ProductUseCase: mockproductUsecase,
		}
		testHelper.PrepareHandler(t,&handler.Controller,req,rec)
		handler.Ctx.Input.SetParam("id",id)
		handler.Prepare()

		handler.GetProductByID()

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

}

