package errors

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

type AppErrorContext struct {
	echo.Context
}

// ErrorResponse generates custom error response of any api
func (c *AppErrorContext) ErrorResponse(err error) error {
	var ge GeneralError
	if isGe := errors.As(err, &ge); isGe {
		httpStatus := GetHTTPStatus(ge.Code())
		er := &errorResponse{
			Code:   ge.Code(),
			Errors: ge.Messages(),
		}
		return c.JSON(httpStatus, er)
	}

	return c.JSON(http.StatusInternalServerError, &errorResponse{
		Code:   InternalServerError,
		Errors: []string{"We are very sorry, internal error occurred."},
	})
}
