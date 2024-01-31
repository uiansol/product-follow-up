package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uiansol/product-follow-up/internal/adapters/api/dto"
	"github.com/uiansol/product-follow-up/internal/adapters/api/mappers"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

type ProductCreateHandler struct {
	productCreateUseCase usecases.IProductCreateUseCase
}

func NewProductCreateHandler(productCreateUseCase usecases.IProductCreateUseCase) *ProductCreateHandler {
	return &ProductCreateHandler{
		productCreateUseCase: productCreateUseCase,
	}
}

func (h *ProductCreateHandler) Handle(c echo.Context) error {
	var product dto.ProductCreateRequest

	if err := c.Bind(&product); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	input := mappers.ProductCreateRequestToProductCreateInput(product)
	output, err := h.productCreateUseCase.Execute(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := mappers.ProductCreateOutputToProductCreateResponse(output)

	return c.JSON(http.StatusOK, response)
}
