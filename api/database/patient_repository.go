package database

import (
	"context"
	"go-technical/api/model"

	"github.com/google/uuid"
)

var patientQuery = PatientQuery{}

type PatientRepository interface {
	CreatePatient(patient model.Patient) (*string, error)
}

func (db *PostgresDatabase) CreatePatient(patient model.Patient) (*string, error) {
	query, _, err := patientQuery.InsertPatient()
	if err != nil {
		return nil, err
	}
	row := db.GetPool().QueryRow(
		context.Background(),
		query,
		uuid.New().String(),
		patient.MemberId,
		patient.ClientId,
		patient.FirstName,
		patient.LastName,
		patient.Gender,
		patient.DateOfBirth,
		patient.PhoneNumber,
		patient.Email,
	)

	var createdId *string
	err = row.Scan(&createdId)
	if err != nil {
		return nil, err
	}

	return createdId, err
}
