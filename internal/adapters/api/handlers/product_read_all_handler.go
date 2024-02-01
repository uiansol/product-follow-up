package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uiansol/product-follow-up/internal/adapters/api/mappers"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

type ProductReadAllHandler struct {
	productReadAllUseCase usecases.IProductReadAllUseCase
}

func NewProductReadAllHandler(productReadAllUseCase usecases.IProductReadAllUseCase) *ProductReadAllHandler {
	return &ProductReadAllHandler{
		productReadAllUseCase: productReadAllUseCase,
	}
}

func (h *ProductReadAllHandler) Handle(c echo.Context) error {
	output, err := h.productReadAllUseCase.Execute()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := mappers.ProductReadAllOutputToProductReadAllResponse(output)

	return c.JSON(http.StatusOK, response)
}
