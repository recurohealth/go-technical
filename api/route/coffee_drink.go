package route

import (
	"go-technical/api/handler"

	"github.com/labstack/echo/v4"
)

func (echoRouter *EchoRouter) AddCoffeeDrinksRoutes(baseGroup *echo.Group) {
	coffeeDrinksGroup := baseGroup.Group(COFFEE_DRINKS_PATH)
	coffeeDrinksGroup.GET("", handler.GetCoffeeDrinks)
	coffeeDrinksGroup.POST("", handler.CreateCoffeeDrink)
}
