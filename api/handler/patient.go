package handler

import (
	"go-technical/api/custom"
	"go-technical/api/database"
	"go-technical/api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePatient(ctx echo.Context) error {
	clientPatientInfo := model.ClientPatientInfo{}
	err := ctx.Bind(&clientPatientInfo)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.NewErrorMessage(
			http.StatusBadRequest,
			"unable to bind request data",
			err,
		))
	}

	db := ctx.Get(custom.DATABASE_MIDDLEWARE_KEY)
	patientRepository := db.(database.PatientRepository)
	patientId, err := patientRepository.CreatePatient(model.Patient{
		ClientId:    clientPatientInfo.ClientId,
		FirstName:   clientPatientInfo.FirstName,
		LastName:    clientPatientInfo.LastName,
		Gender:      clientPatientInfo.Gender,
		DateOfBirth: clientPatientInfo.DateOfBirth,
		PhoneNumber: clientPatientInfo.PhoneNumber,
		Email:       clientPatientInfo.Email,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.NewErrorMessage(
			http.StatusInternalServerError,
			"unable to create patient",
			err,
		))
	}

	addressRepository := db.(database.AddressRepository)
	address := clientPatientInfo.Address
	_, err = addressRepository.CreateAddress(model.Address{
		City:           address.City,
		AddressLineOne: address.AddressLineOne,
		AddressLineTwo: address.AddressLineTwo,
		State:          address.State,
		Zip:            address.Zip,
		Country:        address.Country,
		PatientId:      *patientId,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.NewErrorMessage(
			http.StatusInternalServerError,
			"unable to create address",
			err,
		))
	}

	return ctx.JSON(http.StatusOK, model.Patient{
		Id: *patientId,
	})
}
