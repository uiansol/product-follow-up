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

func TestNewProductCreateHandler(t *testing.T) {
	t.Run("should return new product create handler", func(t *testing.T) {
		useCaseMock := usecases.NewIProductCreateUseCaseMock(t)

		h := NewProductCreateHandler(useCaseMock)

		assert.NotNil(t, h)
	})
}

func TestProductCreateHandle(t *testing.T) {
	t.Run("should process request and return ok with new product id", func(t *testing.T) {
		validBody := dto.ProductCreateRequest{
			Name:     "Test-Name",
			Comments: "Test-Description",
			Link:     "Test-Link",
			Price:    10.5,
		}
		validBodyJSON, _ := json.Marshal(validBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(validBodyJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductCreateUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.ProductCreateOutput{ProductID: "Test-ID"}, nil)

		h := NewProductCreateHandler(useCaseMock)
		h.Handle(c)

		var res map[string]string
		json.NewDecoder(rec.Body).Decode(&res)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "Test-ID", res["product_id"])
	})

	t.Run("should return 400 when cannot bind", func(t *testing.T) {
		invalidBody := "invalid body"

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(invalidBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductCreateUseCaseMock(t)

		h := NewProductCreateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 500 when use case returns error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/product", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := usecases.NewIProductCreateUseCaseMock(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.ProductCreateOutput{}, assert.AnError)

		h := NewProductCreateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
