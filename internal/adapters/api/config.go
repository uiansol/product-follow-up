package api

import (
	"github.com/uiansol/product-follow-up/internal/adapters/api/handlers"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mysql"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
	"gorm.io/gorm"
)

func configRepositories(db *gorm.DB) *AppRepositories {
	productRepository := mysql.NewProductRepository(db)

	return &AppRepositories{
		productRepository: productRepository,
	}
}

func configUseCases(repositories *AppRepositories) *AppUseCases {
	productCreateUseCase := usecases.NewProductCreateUseCase(repositories.productRepository)

	return &AppUseCases{
		productCreateUseCase: productCreateUseCase,
	}
}

func configHandlers(usecases *AppUseCases) *AppHandlers {
	pingHandler := handlers.NewPingHandler()

	productCreateHandler := handlers.NewProductCreateHandler(usecases.productCreateUseCase)

	return &AppHandlers{
		pingHandler:    pingHandler,
		productHandler: productCreateHandler,
	}
}
