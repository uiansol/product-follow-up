package api

import (
	"database/sql"

	"github.com/uiansol/product-follow-up/internal/adapters/api/handlers"
	"github.com/uiansol/product-follow-up/internal/adapters/db/mysql"
	"github.com/uiansol/product-follow-up/internal/application/usecases"
)

func configHandlers(usecases *AppUseCases) *AppHandlers {
	pingHandler := handlers.NewPingHandler()

	productCreateHandler := handlers.NewProductCreateHandler(usecases.productCreateUseCase)

	return &AppHandlers{
		pingHandler:    pingHandler,
		productHandler: productCreateHandler,
	}
}

func configRepositories(db *sql.DB) *AppRepositories {
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
