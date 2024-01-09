package database

import (
	"context"
	"go-technical/api/model"

	"github.com/google/uuid"
)

var addressQuery = AddressQuery{}

type AddressRepository interface {
	CreateAddress(address model.Address) (*string, error)
}

func (db *PostgresDatabase) CreateAddress(address model.Address) (*string, error) {
	query, _, err := addressQuery.InsertAddress()
	if err != nil {
		return nil, err
	}
	row := db.GetPool().QueryRow(
		context.Background(),
		query,
		uuid.New().String(),
		address.City,
		address.AddressLineOne,
		address.AddressLineTwo,
		address.State,
		address.Zip,
		address.Country,
		address.PatientId,
	)

	var createdId *string
	err = row.Scan(&createdId)
	if err != nil {
		return nil, err
	}

	return createdId, err
}
