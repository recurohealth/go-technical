package database

import (
	"context"
	"go-technical/api/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

var coffeeDrinksQuery = CoffeeDrinkQuery{}

type CoffeeDrinksRepository interface {
	GetCoffeeDrinks() (*model.CoffeeDrinks, error)
	GetCoffeeDrinksByTemperatureStyle(temperatureStyle model.TemperatureStyle) (*model.CoffeeDrinks, error)
	CreateCoffeeDrink(coffeeDrink model.CoffeeDrink) (*string, error)
}

func (db *PostgresDatabase) GetCoffeeDrinks() (*model.CoffeeDrinks, error) {
	coffeeDrinks := &model.CoffeeDrinks{}
	query, _, err := coffeeDrinksQuery.SelectCoffeeDrinks()
	if err != nil {
		return nil, err
	}
	err = pgxscan.Select(
		context.Background(),
		db.GetPool(),
		coffeeDrinks,
		query,
	)
	if err != nil {
		return nil, err
	}

	return coffeeDrinks, err
}

func (db *PostgresDatabase) GetCoffeeDrinksByTemperatureStyle(temperatureStyle model.TemperatureStyle) (*model.CoffeeDrinks, error) {
	coffeeDrinks := &model.CoffeeDrinks{}
	query, _, err := coffeeDrinksQuery.SelectCoffeeDrinksByTemperatureStyle()
	if err != nil {
		return nil, err
	}
	err = pgxscan.Select(
		context.Background(),
		db.GetPool(),
		coffeeDrinks,
		query,
		temperatureStyle,
	)
	if err != nil {
		return nil, err
	}

	return coffeeDrinks, err
}

func (db *PostgresDatabase) CreateCoffeeDrink(coffeeDrink model.CoffeeDrink) (*string, error) {
	query, _, err := coffeeDrinksQuery.InsertCoffeeDrink()
	if err != nil {
		return nil, err
	}
	row := db.GetPool().QueryRow(
		context.Background(),
		query,
		uuid.New().String(),
		coffeeDrink.Name,
		coffeeDrink.Description,
		coffeeDrink.Origin,
		coffeeDrink.TemperatureStyle,
	)

	var createdId *string
	err = row.Scan(&createdId)
	if err != nil {
		return nil, err
	}

	return createdId, err
}
