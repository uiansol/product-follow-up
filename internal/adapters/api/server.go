package api

import (
	"github.com/labstack/echo/v4"
	"github.com/uiansol/product-follow-up/internal/adapters/api/handlers"
)

type RestServer struct {
	router     *echo.Echo
	appHandler *AppHandlers
}

type AppHandlers struct {
	pingHandler    *handlers.PingHandler
	productHandler *handlers.ProductCreateHandler
}

func NewRestService(router *echo.Echo, appHandler *AppHandlers) *RestServer {
	return &RestServer{
		router:     router,
		appHandler: appHandler,
	}
}

func SetUpServer() {
	router := echo.New()
	appHandler := configHandlers()

	server := NewRestService(router, appHandler)
	server.SetUpRoutes()

	server.router.Logger.Fatal(router.Start(":8080"))
}
