package api

import (
	"github.com/labstack/echo/v4"
)

func (s *RestServer) SetUpRoutes() {
	v1 := s.router.Group("/v1")
	s.PingRoutes(v1)

	v2 := s.router.Group("/v2")
	s.ProductRoutes(v2)
}

func (s *RestServer) PingRoutes(v1 *echo.Group) {
	v1.GET("/ping", func(c echo.Context) error {
		return s.appHandler.pingHandler.Handle(c)
	})
}

func (s *RestServer) ProductRoutes(v2 *echo.Group) {
	product := v2.Group("/product")

	product.POST("/", func(c echo.Context) error {
		return s.appHandler.productCreateHandler.Handle(c)
	})
	product.GET("/", func(c echo.Context) error {
		return s.appHandler.productReadAllHandler.Handle(c)
	})
	product.GET("/:id", func(c echo.Context) error {
		return s.appHandler.productReadHandler.Handle(c)
	})
	product.PUT("/:id", func(c echo.Context) error {
		return s.appHandler.productUpdateHandler.Handle(c)
	})
	product.DELETE("/:id", func(c echo.Context) error {
		return s.appHandler.productDeleteHandler.Handle(c)
	})
}
