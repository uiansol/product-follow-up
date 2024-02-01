package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/adapters/api/mappers"
	"github.com/uiansol/product-follow-up/internal/application/apperr"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

type ProductDeleteHandler struct {
	productDeleteUseCase usecases.IProductDeleteUseCase
}

func NewProductDeleteHandler(productDeleteUseCase usecases.IProductDeleteUseCase) *ProductDeleteHandler {
	return &ProductDeleteHandler{
		productDeleteUseCase: productDeleteUseCase,
	}
}

func (h *ProductDeleteHandler) Handle(c echo.Context) error {
	var request dto.ProductIDRequest

	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	input := mappers.ProductIDRequestToProductDeleteInput(request)
	err := h.productDeleteUseCase.Execute(input)
	if err != nil {
		if err.Error() == apperr.ErrNotFound {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
