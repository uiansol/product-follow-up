package api

import (
	"github.com/labstack/echo/v4"
)

func (s *RestServer) SetUpRoutes() {
	s.PingRoutes()
	s.ProductRoutes()
}

func (s *RestServer) PingRoutes() {
	s.router.GET("/ping", func(c echo.Context) error {
		return s.appHandler.pingHandler.Handle(c)
	})
}

func (s *RestServer) ProductRoutes() {
	s.router.POST("/product", func(c echo.Context) error {
		return s.appHandler.productCreateHandler.Handle(c)
	})
	s.router.GET("/product", func(c echo.Context) error {
		return s.appHandler.productReadAllHandler.Handle(c)
	})
	s.router.GET("/product/:id", func(c echo.Context) error {
		return s.appHandler.productReadHandler.Handle(c)
	})
	s.router.PUT("/product/:id", func(c echo.Context) error {
		return s.appHandler.productUpdateHandler.Handle(c)
	})
	s.router.DELETE("/product/:id", func(c echo.Context) error {
		return s.appHandler.productDeleteHandler.Handle(c)
	})
}
