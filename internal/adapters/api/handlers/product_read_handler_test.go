package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/application/apperr"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

func TestNewProductReadHandler(t *testing.T) {
	t.Run("should return new product read handler", func(t *testing.T) {
		useCaseMock := usecases.NewIProductReadUseCaseMock(t)

		h := NewProductReadHandler(useCaseMock)

		assert.NotNil(t, h)
	})
}

func TestProductReadHandle(t *testing.T) {
	t.Run("should process request and return ok with product", func(t *testing.T) {
		date, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
		product := usecases.ProductReadOutput{
			ID:        "Test-ID",
			Name:      "Test",
			Comments:  "Test comments",
			Link:      "http://test.com",
			Price:     10.5,
			PriceDate: date,
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v2/product/Test-ID", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductReadUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(product, nil)

		h := NewProductReadHandler(useCaseMock)
		h.Handle(c)

		var res dto.ProductReadResponse
		json.NewDecoder(rec.Body).Decode(&res)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, product.ID, res.ID)
		assert.Equal(t, product.Name, res.Name)
		assert.Equal(t, product.Comments, res.Comments)
		assert.Equal(t, product.Link, res.Link)
		assert.Equal(t, 10.5, res.Price)
		assert.Equal(t, product.PriceDate, res.PriceDate)
	})

	t.Run("should return 404 when use case returns error not found", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v2/product/Test-ID", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductReadUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.ProductReadOutput{}, errors.New(apperr.ErrNotFound))

		h := NewProductReadHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("should return 500 when use case returns error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v2/product/Test-ID", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductReadUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.ProductReadOutput{}, assert.AnError)

		h := NewProductReadHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
