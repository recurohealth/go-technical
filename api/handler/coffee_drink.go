package handler

import (
	"net/http"

	"go-template/api/custom"
	"go-template/api/database"
	"go-template/api/model"

	"github.com/labstack/echo/v4"
)

func CreateCoffeeDrink(ctx echo.Context) error {
	coffeeDrink := model.CoffeeDrink{}
	err := ctx.Bind(&coffeeDrink)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"unable to bind request data",
			err,
		))
	}
	err = model.Validator.Struct(&coffeeDrink)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"invalid coffee drink request data",
			err,
		))
	}

	repository := ctx.Get(custom.DATABASE_MIDDLEWARE_KEY).(database.CoffeeDrinksRepository)
	coffeeDrinkId, err := repository.CreateCoffeeDrink(coffeeDrink)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.NewErrorMessage(
			http.StatusInternalServerError,
			"unable to create coffee drink",
			err,
		))
	}

	return ctx.JSON(http.StatusCreated, coffeeDrinkId)
}

func GetCoffeeDrinks(ctx echo.Context) error {
	getCoffeeDrinks := model.GetCoffeeDrinks{}
	err := ctx.Bind(&getCoffeeDrinks)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"unable to bind request data",
			err,
		))
	}
	err = model.Validator.Struct(&getCoffeeDrinks)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"invalid get coffee drinks request data",
			err,
		))
	}

	repository := ctx.Get(custom.DATABASE_MIDDLEWARE_KEY).(database.CoffeeDrinksRepository)

	if getCoffeeDrinks.TemperatureStyle != nil {
		coffeeDrinksByTemperature, err := repository.GetCoffeeDrinksByTemperatureStyle(*getCoffeeDrinks.TemperatureStyle)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, model.NewErrorMessage(
				http.StatusInternalServerError,
				"unable to get coffee drinks by temperature style",
				err,
			))
		}

		return ctx.JSON(http.StatusOK, coffeeDrinksByTemperature)
	}

	coffeeDrinks, err := repository.GetCoffeeDrinks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.NewErrorMessage(
			http.StatusInternalServerError,
			"unable to get coffee drinks",
			err,
		))
	}

	return ctx.JSON(http.StatusOK, coffeeDrinks)
}
