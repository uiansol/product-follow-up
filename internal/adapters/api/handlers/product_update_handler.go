package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/adapters/api/mappers"
	"github.com/uiansol/product-follow-up/internal/application/apperr"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

type ProductUpdateHandler struct {
	productUpdateUseCase usecases.IProductUpdateUseCase
}

func NewProductUpdateHandler(productUpdateUseCase usecases.IProductUpdateUseCase) *ProductUpdateHandler {
	return &ProductUpdateHandler{
		productUpdateUseCase: productUpdateUseCase,
	}
}

func (h *ProductUpdateHandler) Handle(c echo.Context) error {
	var product dto.ProductUpdateRequest

	if err := c.Bind(&product); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	input := mappers.ProductUpdateRequestToProductUpdateInput(product)
	err := h.productUpdateUseCase.Execute(input)
	if err != nil {
		if err.Error() == apperr.ErrNotFound {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
