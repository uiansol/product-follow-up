package api

import (
	"github.com/uiansol/product-follow-up/internal/adapters/api/handlers"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mysql"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

func configHandlers() *AppHandlers {
	pingHandler := handlers.NewPingHandler()

	productRepository := mysql.NewProductRepository()

	productCreateUseCase := usecases.NewProductCreateUseCase(productRepository)
	productCreateHandler := handlers.NewProductCreateHandler(&productCreateUseCase)

	return &AppHandlers{
		pingHandler:    pingHandler,
		productHandler: productCreateHandler,
	}
}
