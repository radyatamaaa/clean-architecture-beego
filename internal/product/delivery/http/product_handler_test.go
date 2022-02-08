package http_test

import (
	"clean-architecture-beego/internal/domain"
	productHttpHandler "clean-architecture-beego/internal/product/delivery/http"
	_productUsecaseMock "clean-architecture-beego/internal/product/mocks"
	"context"
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

}
