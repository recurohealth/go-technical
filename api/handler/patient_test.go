package handler

import (
	"bytes"
	"encoding/json"
	"go-technical/api/custom"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
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
		databaseStub := DatabaseStub{}
		ctx.Set(custom.DATABASE_MIDDLEWARE_KEY, &databaseStub)

		_ = CreatePatient(ctx)
	})
}

type DatabaseStub struct{}
