package api

import "github.com/uiansol/product-follow-up/internal/adapters/api/handlers"

func configHandlers() *AppHandlers {
	pingHandler := handlers.NewPingHandler()

	return &AppHandlers{
		pingHandler: pingHandler,
	}
}
