package database

import (
	"github.com/doug-martin/goqu/v9"
)

const PATIENT_TABLE = "patient"

type PatientQuery struct{}

func (PatientQuery) InsertPatient() (query string, args []interface{}, err error) {
	return postgresDialect.Insert(PATIENT_TABLE).
		Prepared(true).
		Rows(
			goqu.Record{
				"id": goqu.L("$1"),
				// TODO: add other fields
			},
		).
		Returning("id").
		ToSQL()
}
