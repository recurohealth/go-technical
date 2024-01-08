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
		data, _ := json.Marshal(map[string]interface{}{})
		requestBody := bytes.NewReader(data)
		request := httptest.NewRequest(http.MethodPost, "/patients", requestBody)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		ctx := e.NewContext(request, recorder)
		databaseCreatePatientStub := DatabaseCreatePatientStub{}
		ctx.Set(custom.DATABASE_MIDDLEWARE_KEY, &databaseCreatePatientStub)

		_ = CreatePatient(ctx)

		assert.Equal(t, "expected", "actual")
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
