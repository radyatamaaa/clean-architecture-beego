package http_test

import (
	"clean-architecture-beego/internal/domain"
	productHttpHandler "clean-architecture-beego/internal/product/delivery/http"
	_productUsecaseMock "clean-architecture-beego/internal/product/mocks"
	"context"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
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

//func Init() {
//	_, file, _, _ := runtime.Caller(0)
//	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
//	beego.TestBeegoInit(apppath)
//}

//func prepareController (c * beego.Controller,request *http.Request ,response *beegoContext.Response) {
//	c.Ctx = &beegoContext.Context {
//		Request: request,
//		ResponseWriter: response,
//	}
//	c.Ctx.Output = &beegoContext.BeegoOutput {Context: c.Ctx}
//	//c.Ctx.Input = &beegoContext.BeegoInput {Request: c.Ctx.Request}
//	//
//	//globalSessions, _:= session.NewManager ("memory", `{" cookieName ":" gosessionid "," gclifetime ": 10}`)
//	//c.Ctx.Request.Header = http.Header {}
//	//c.Ctx.Request.AddCookie (& http.Cookie {Name: "gosessionid", Value: "test"})
//	//c.CruSession = globalSessions.SessionRegenerateId (c.Ctx.ResponseWriter, c.Ctx.Request)
//	//c.Data = map [interface {}] interface {} {}
//}
//type fakeResponseWriter struct {}
//
//func (f * fakeResponseWriter) Header () http.Header {
//	return http.Header {}
//}
//func (f * fakeResponseWriter) Write (b [] byte) (int, error) {
//	return 0, nil
//}
//func (f * fakeResponseWriter) WriteHeader (n int) {}
func PrepareHandler(t *testing.T,handler * beego.Controller,request *http.Request,response http.ResponseWriter) {
	err := request.ParseForm()
	assert.NoError(t, err)

	handler.Ctx = &beegoContext.Context{
		Request:        request,
		ResponseWriter: &beegoContext.Response{
			ResponseWriter: response,
			Started:        false,
			Status:         0,
			Elapsed:        0,
		},
	}
	body, _ := ioutil.ReadAll(handler.Ctx.Request.Body)
	handler.Ctx.Input = &beegoContext.BeegoInput{
		Context:       handler.Ctx,
		CruSession:    nil,
		RequestBody:   body,
		RunMethod:     "",
		RunController: nil,
	}
	handler.Ctx.Output = &beegoContext.BeegoOutput{
		Context:    handler.Ctx,
	}
	handler.Data = map[interface{}]interface{} {}

}
func TestProductHandler_GetProducts(t *testing.T) {
	//requestMock
	offset := 0
	limit := 10

	//resultMock
	mockProduct := []domain.Product{}
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
		PrepareHandler(t,&handler.Controller,req,rec)
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
		PrepareHandler(t,&handler.Controller,req,rec)
		handler.Prepare()

		handler.GetProducts()

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockproductUsecase.AssertExpectations(t)

	})

}
