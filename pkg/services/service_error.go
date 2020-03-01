package services

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServiceError interface {
	Error() string
	ToHTTPError() *echo.HTTPError
}

type serviceErrorImpl struct {
	httpCode int
	message  string
	cause    error
}

func (s *serviceErrorImpl) Error() string {
	return fmt.Sprintf("%s. Server error: '%s'", s.message, s.cause.Error())
}

func (s serviceErrorImpl) ToHTTPError() *echo.HTTPError {
	return echo.NewHTTPError(s.httpCode, s.Error())
}

func NewGenericServiceError(message string, cause error) ServiceError {
	return &serviceErrorImpl{
		message:  message,
		httpCode: http.StatusInternalServerError,
		cause:    cause,
	}
}

func NewNotFoundError(message string, cause error) ServiceError {
	return &serviceErrorImpl{
		httpCode: http.StatusNotFound,
		message:  message,
		cause:    cause,
	}
}

func NewBadReadingTimeError(message string, cause error) ServiceError {
	return &serviceErrorImpl{
		httpCode: http.StatusBadRequest,
		message:  message,
		cause:    cause,
	}
}
