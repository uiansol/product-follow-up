package api

import (
	"github.com/labstack/echo/v4"
)

func (s *RestServer) SetUpRoutes() {
	s.PingRoutes()
}

func (s *RestServer) PingRoutes() {
	s.router.GET("/ping", func(c echo.Context) error {
		return s.appHandler.pingHandler.Handle(c)
	})
}
