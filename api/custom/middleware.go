package custom

import (
	"go-technical/api/database"

	"github.com/labstack/echo/v4"
)

const DATABASE_MIDDLEWARE_KEY = "db"

func DatabaseMiddleware(db database.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set(DATABASE_MIDDLEWARE_KEY, db)

			return next(ctx)
		}
	}
}
