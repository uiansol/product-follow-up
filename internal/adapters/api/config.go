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
	productReadUseCase := usecases.NewProductReadUseCase(repositories.productRepository)
	productUpdateUseCase := usecases.NewProductUpdateUseCase(repositories.productRepository)
	productDeleteUseCase := usecases.NewProductDeleteUseCase(repositories.productRepository)

	return &AppUseCases{
		productCreateUseCase: productCreateUseCase,
		productReadUseCase:   productReadUseCase,
		productUpdateUseCase: productUpdateUseCase,
		productDeleteUseCase: productDeleteUseCase,
	}
}

func configHandlers(usecases *AppUseCases) *AppHandlers {
	pingHandler := handlers.NewPingHandler()

	productCreateHandler := handlers.NewProductCreateHandler(usecases.productCreateUseCase)
	productReadHandler := handlers.NewProductReadHandler(usecases.productReadUseCase)
	productUpdateHandler := handlers.NewProductUpdateHandler(usecases.productUpdateUseCase)
	productDeleteHandler := handlers.NewProductDeleteHandler(usecases.productDeleteUseCase)

	return &AppHandlers{
		pingHandler:          pingHandler,
		productCreateHandler: productCreateHandler,
		productReadHandler:   productReadHandler,
		productUpdateHandler: productUpdateHandler,
		productDeleteHandler: productDeleteHandler,
	}
}
