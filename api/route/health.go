package route

import (
	"go-template/api/handler"

	"github.com/labstack/echo/v4"
)

func (echoRouter *EchoRouter) AddHealthRoute(baseGroup *echo.Group) {
	healthRoute := baseGroup.Group(HEALTH_ROUTE)
	healthRoute.GET("", handler.Health)
}
