package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePatient(ctx echo.Context) error {
	// patient := model.Patient{}
	// err := ctx.Bind(&patient)
	// if err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, model.NewErrorMessage(
	// 		http.StatusBadRequest,
	// 		"unable to bind request data",
	// 		err,
	// 	))
	// }

	// patientRepository := ctx.Get(custom.DATABASE_MIDDLEWARE_KEY).(database.PatientRepository)

	return ctx.JSON(http.StatusOK, "some id")
}
