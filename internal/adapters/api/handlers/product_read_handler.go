package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/adapters/api/mappers"
	"github.com/uiansol/product-follow-up/internal/application/apperr"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

type ProductReadHandler struct {
	productReadUseCase usecases.IProductReadUseCase
}

func NewProductReadHandler(productReadUseCase usecases.IProductReadUseCase) *ProductReadHandler {
	return &ProductReadHandler{
		productReadUseCase: productReadUseCase,
	}
}

func (h *ProductReadHandler) Handle(c echo.Context) error {
	var request dto.ProductIDRequest

	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	input := mappers.ProductIDRequestToProductReadInput(request)
	output, err := h.productReadUseCase.Execute(input)
	if err != nil {
		if err.Error() == apperr.ErrNotFound {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := mappers.ProductReadOutputToProductReadResponse(output)

	return c.JSON(http.StatusOK, response)
}
