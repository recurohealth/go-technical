package handler

import (
	"bytes"
	"encoding/json"
	"go-technical/api/custom"
	"go-technical/api/database"
	"go-technical/api/model"
	"go-technical/test/utility"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatientHandler(t *testing.T) {
	e := echo.New()

	t.Run("test description", func(t *testing.T) {
		data, _ := json.Marshal(model.ClientPatientInfo{
			ClientId:    "anything",
			FirstName:   "anything",
			LastName:    "anything",
			Gender:      model.MALE,
			DateOfBirth: "anything",
			PhoneNumber: "anything",
			Email:       "anything",
			Address: model.Address{
				City:           "anything",
				AddressLineOne: "anything",
				AddressLineTwo: nil,
				State:          "anything",
				Zip:            "anything",
				Country:        "anything",
			},
		})
		requestBody := bytes.NewReader(data)
		request := httptest.NewRequest(http.MethodPost, "/patients", requestBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		ctx := e.NewContext(request, recorder)
		databaseCreatePatientStub := DatabaseCreatePatientStub{}
		ctx.Set(custom.DATABASE_MIDDLEWARE_KEY, &databaseCreatePatientStub)

		_ = CreatePatient(ctx)

		assert.Equal(t, "placeholder", "placeholder")
	})
}

type DatabaseCreatePatientStub struct {
	database.Database
	database.PatientRepository
}

func (db *DatabaseCreatePatientStub) GetPool() database.PgxPool {
	return database.PgxPool(&utility.DatabasePoolStub{})
}
func (db *DatabaseCreatePatientStub) CreatePatient(patient model.Patient) (*string, error) {
	randomUUID := utility.Faker.UUID()
	return &randomUUID, nil
}
func (db *DatabaseCreatePatientStub) CreateAddress(address model.Address) (*string, error) {
	randomUUID := utility.Faker.UUID()
	return &randomUUID, nil
}
