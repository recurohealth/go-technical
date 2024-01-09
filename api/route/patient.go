package route

import (
	"go-technical/api/handler"

	"github.com/labstack/echo/v4"
)

func (echoRouter *EchoRouter) AddPatientRoutes(baseGroup *echo.Group) {
	patientsGroup := baseGroup.Group(PATIENTS_PATH)
	patientsGroup.POST("", handler.CreatePatient)
}
