package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

func TestNewProductReadAllHandler(t *testing.T) {
	t.Run("should return new product read all handler", func(t *testing.T) {
		useCaseMock := usecases.NewIProductReadAllUseCaseMock(t)

		h := NewProductReadAllHandler(useCaseMock)

		assert.NotNil(t, h)
	})
}

func TestProductReadAllHandle(t *testing.T) {
	t.Run("should process request and return ok with product", func(t *testing.T) {
		date, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

		var products []*usecases.ProductReadAllOutput
		products = append(products, &usecases.ProductReadAllOutput{
			ID:        "Test-ID",
			Name:      "Test",
			Comments:  "Test comments",
			Link:      "http://test.com",
			Price:     10.5,
			PriceDate: date,
		})
		products = append(products, &usecases.ProductReadAllOutput{
			ID:        "Test-ID 2",
			Name:      "Test 2",
			Comments:  "Test comments 2",
			Link:      "http://test.com",
			Price:     13.5,
			PriceDate: date,
		})

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/product", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductReadAllUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(products, nil)

		h := NewProductReadAllHandler(useCaseMock)
		h.Handle(c)

		var res dto.ProductReadAllResponse
		json.NewDecoder(rec.Body).Decode(&res)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 2, len(res.Products))

		assert.Equal(t, "Test-ID", res.Products[0].ID)
		assert.Equal(t, "Test", res.Products[0].Name)
		assert.Equal(t, "Test comments", res.Products[0].Comments)
		assert.Equal(t, "http://test.com", res.Products[0].Link)
		assert.Equal(t, 10.5, res.Products[0].Price)
		assert.Equal(t, date, res.Products[0].PriceDate)

		assert.Equal(t, "Test-ID 2", res.Products[1].ID)
		assert.Equal(t, "Test 2", res.Products[1].Name)
		assert.Equal(t, "Test comments 2", res.Products[1].Comments)
		assert.Equal(t, "http://test.com", res.Products[1].Link)
		assert.Equal(t, 13.5, res.Products[1].Price)
		assert.Equal(t, date, res.Products[1].PriceDate)
	})

	t.Run("should return 500 when use case returns error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/product/Test-ID", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductReadAllUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(nil, assert.AnError)

		h := NewProductReadAllHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
