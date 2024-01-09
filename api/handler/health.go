package handler

import (
	"go-technical/api/custom"
	"go-technical/api/database"
	"go-technical/api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Health(ctx echo.Context) error {
	db := ctx.Get(custom.DATABASE_MIDDLEWARE_KEY).(database.Database)
	err := db.GetPool().Ping(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusServiceUnavailable, model.NewHealthMessage(
			http.StatusServiceUnavailable,
			"Unable to connect to database",
		))
	}

	return ctx.JSON(http.StatusOK, model.NewHealthMessage(
		http.StatusOK,
		"Service is up and running",
	))
}
