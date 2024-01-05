package database

import (
	"github.com/doug-martin/goqu/v9"
)

const COFFEE_DRINKS_TABLE = "coffee_drink"

type CoffeeDrinkQuery struct{}

func (CoffeeDrinkQuery) SelectCoffeeDrinks() (query string, args []interface{}, err error) {
	return postgresDialect.From(COFFEE_DRINKS_TABLE).ToSQL()
}

func (CoffeeDrinkQuery) SelectCoffeeDrinksByTemperatureStyle() (query string, args []interface{}, err error) {
	return postgresDialect.From(COFFEE_DRINKS_TABLE).
		Prepared(true).
		Where(goqu.Ex{"temperature_style": goqu.Op{"eq": "$1"}}).
		ToSQL()
}

func (CoffeeDrinkQuery) InsertCoffeeDrink() (query string, args []interface{}, err error) {
	return postgresDialect.Insert(COFFEE_DRINKS_TABLE).
		Prepared(true).
		Rows(
			goqu.Record{
				"id":                goqu.L("$1"),
				"name":              goqu.L("$2"),
				"description":       goqu.L("$3"),
				"origin":            goqu.L("$4"),
				"temperature_style": goqu.L("$5"),
			},
		).
		Returning("id").
		ToSQL()
}
