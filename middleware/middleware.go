package middleware

import (
	"github.com/labstack/echo/v4"
	"clean-architecture-beego/helper"
	"clean-architecture-beego/helper/logger"
	models2 "clean-architecture-beego/helper/models"
	"net/http"
	"strconv"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("Access-Control-Allow-Headers", "*")
		c.Request().Header.Set("Access-Control-Allow-Origin", "*")
		c.Request().Header.Set("Access-Control-Allow-Methods", "*")

		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Headers", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "*")
		return next(c)
	}
}

// LOG will handle the LOG middleware
func (m *GoMiddleware) Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l := logger.L
		l.Info("Accepted")

		next(c)

		l.Info("[" + strconv.Itoa(c.Response().Status) + "] " + "[" + c.Request().Method + "] " + c.Request().Host + c.Request().URL.String())

		l.Info("Closing")
		return nil
	}
}

// CORSValidationGlobalResponse will handle the CORSValidationGlobalResponse middleware
func (m *GoMiddleware) CORSValidationGlobalResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" &&
			c.Request().Header.Get(echo.HeaderContentType) == echo.MIMEApplicationJSON {
			c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			response := new(models2.Response)
			response.MappingResponseError(http.StatusBadRequest, "Invalid "+
				echo.HeaderContentType+" "+echo.MIMEApplicationJSON, nil)
			c.JSON(response.StatusCode, response)
			return nil
		}


		err := next(c)

		if err != nil {
			allowedMethod := []string{"GET", "POST"}
			if !helper.InArray(c.Request().Method, allowedMethod) {
				response := new(models2.Response)
				response.MappingResponseError(http.StatusBadRequest, "Invalid Method", nil)
				c.JSON(response.StatusCode, response)
				return err
			} else if err.Error() == "code=404, message=Not Found" {
				response := new(models2.Response)
				response.MappingResponseError(http.StatusNotFound, "Routes Not Found", nil)

				c.JSON(response.StatusCode, response)
				return err
			} else if err.Error() == "code=405, message=Method Not Allowed" {
				response := new(models2.Response)
				response.MappingResponseError(http.StatusMethodNotAllowed, "Method Not Allowed", nil)

				c.JSON(response.StatusCode, response)
				return err
			}
		}
		return nil
	}
}

// InitMiddleware intialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
