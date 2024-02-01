package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

func TestNewProductUpdateHandler(t *testing.T) {
	t.Run("should return new product Update handler", func(t *testing.T) {
		useCaseMock := usecases.NewIProductUpdateUseCaseMock(t)

		h := NewProductUpdateHandler(useCaseMock)

		assert.NotNil(t, h)
	})
}

func TestProductUpdateHandle(t *testing.T) {
	t.Run("should process request and return ok", func(t *testing.T) {
		validBody := dto.ProductUpdateRequest{
			Name:     "Test-Name",
			Comments: "Test-Description",
			Link:     "Test-Link",
			Price:    10.5,
		}
		validBodyJSON, _ := json.Marshal(validBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/product/Test-ID", strings.NewReader(string(validBodyJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductUpdateUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(nil)

		h := NewProductUpdateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 400 when cannot bind", func(t *testing.T) {
		invalidBody := "invalid body"

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/product/Test-ID", strings.NewReader(invalidBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductUpdateUseCaseMock(t)

		h := NewProductUpdateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 500 when use case returns error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/product/Test-ID", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductUpdateUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(assert.AnError)

		h := NewProductUpdateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
