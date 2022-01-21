package http_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"clean-architecture-beego/helper/logger"
	models2 "clean-architecture-beego/helper/models"
	"clean-architecture-beego/models"
	sampleModuleHttp "clean-architecture-beego/sample_module/delivery/http"
	sampleModuleHttpHandler "clean-architecture-beego/sample_module/delivery/http"
	_sampleModuleUsecaseMock "clean-architecture-beego/sample_module/mocks"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var (
	l             = logger.L
	groupRoutes   = "/api"
	Documentation = "/"
	List          = groupRoutes + "/sample-module"
)

func TestSampleModuleHandler_LoadHandler(t *testing.T) {
	mockSampleModuleUsecase := new(_sampleModuleUsecaseMock.Usecase)
	e := echo.New()
	sampleModuleHttpHandler.NewsampleModuleHandler(e, mockSampleModuleUsecase, l)

	mockSampleModuleUsecase.AssertExpectations(t)
}
func TestSampleModuleHandler_List(t *testing.T) {

	//requestMock
	limit := 5
	page := 1

	//resultMock
	mockSampleModulePaginationDto := models.SampleModulePaginationDto{}
	mockSampleModulePaginationDto = mockSampleModulePaginationDto.MappingExpampleData()

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockSampleModuleUsecase := new(_sampleModuleUsecaseMock.Usecase)

		mockSampleModuleUsecase.On("GetList", mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&mockSampleModulePaginationDto, models2.GeneralSuccess, nil)

		//j, err := json.Marshal(mockClientDetailRequest)
		e := echo.New()
		queryparam := "?page=" + strconv.Itoa(page) + "&limit=" + strconv.Itoa(limit)
		req, err := http.NewRequest(echo.GET, List+queryparam, strings.NewReader(""))
		//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(List)
		c.Request().ParseForm()
		handler := sampleModuleHttp.SampleModuleHandler{
			SampleModuleUsecase: mockSampleModuleUsecase,
			Log:                 l,
		}
		err = handler.List(c)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)
		mockSampleModuleUsecase.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		//usecaseMock
		mockSampleModuleUsecase := new(_sampleModuleUsecaseMock.Usecase)

		mockSampleModuleUsecase.On("GetList", mock.Anything,
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, models2.ErrGeneralMessage.Error(), models2.ErrInternalServerError)

		//j, err := json.Marshal(mockClientDetailRequest)
		e := echo.New()
		queryparam := "?page=" + strconv.Itoa(page) + "&limit=" + strconv.Itoa(limit)
		req, err := http.NewRequest(echo.GET, List+queryparam, strings.NewReader(""))
		//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(List)
		c.Request().ParseForm()
		handler := sampleModuleHttp.SampleModuleHandler{
			SampleModuleUsecase: mockSampleModuleUsecase,
			Log:                 l,
		}
		err = handler.List(c)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockSampleModuleUsecase.AssertExpectations(t)
	})
}

func TestSampleModuleHandler_Documentation(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		//usecaseMock
		mockSampleModuleUsecase := new(_sampleModuleUsecaseMock.Usecase)

		//j, err := json.Marshal(mockClientDetailRequest)
		e := echo.New()
		req, err := http.NewRequest(echo.GET, Documentation, strings.NewReader(""))
		//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(List)
		c.Request().ParseForm()
		handler := sampleModuleHttp.SampleModuleHandler{
			SampleModuleUsecase: mockSampleModuleUsecase,
			Log:                 l,
		}
		err = handler.Documentation(c)
		require.NoError(t, err)

		assert.Equal(t, 301, rec.Code)
	})

}
