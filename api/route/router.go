package route

import (
	"fmt"
	"go-technical/api/custom"
	"go-technical/api/database"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

const BASE_PATH = "api/v1"

type EchoRouter struct {
	*echo.Echo
}

func NewEchoRouter(db database.Database) *EchoRouter {
	e := echo.New()
	e.Use(custom.DatabaseMiddleware(db))
	e.Use(middleware.BodyDump(func(ctx echo.Context, requestBody, responseBody []byte) {
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		statusCode := ctx.Response().Status
		baseLogInfo := logger.Info().
			Str("uri", ctx.Request().RequestURI).
			Str("method", ctx.Request().Method).
			Str("status", fmt.Sprintf("%v %v", statusCode, http.StatusText(statusCode)))

		if statusCode == http.StatusOK || statusCode == http.StatusCreated {
			baseLogInfo.Msg("Success")
		} else {
			baseLogInfo.Msg(string(responseBody))
		}
	}))

	return &EchoRouter{e}
}

func (echoRouter *EchoRouter) RegisterRoutes() {
	baseGroup := echoRouter.Group(BASE_PATH)
	echoRouter.AddCoffeeDrinksRoutes(baseGroup)
	echoRouter.AddHealthRoute(baseGroup)
}
