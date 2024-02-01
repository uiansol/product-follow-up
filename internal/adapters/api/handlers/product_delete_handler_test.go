package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/application/apperr"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

func TestNewProductDeleteHandler(t *testing.T) {
	t.Run("should return new product delete handler", func(t *testing.T) {
		useCaseMock := usecases.NewIProductDeleteUseCaseMock(t)

		h := NewProductDeleteHandler(useCaseMock)

		assert.NotNil(t, h)
	})
}

func TestProductDeleteHandle(t *testing.T) {
	t.Run("should process request and return ok", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/product/Test-ID", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductDeleteUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(nil)

		h := NewProductDeleteHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 404 when use case returns error not found", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/product/Test-ID", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductDeleteUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(errors.New(apperr.ErrNotFound))

		h := NewProductDeleteHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("should return 500 when use case returns error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/product/Test-ID", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductDeleteUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(errors.New("error"))

		h := NewProductDeleteHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
