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
				"id":            goqu.L("$1"),
				"client_id":     goqu.L("$2"),
				"first_name":    goqu.L("$3"),
				"last_name":     goqu.L("$4"),
				"gender":        goqu.L("$5"),
				"date_of_birth": goqu.L("$6"),
				"phone_number":  goqu.L("$7"),
				"email":         goqu.L("$8"),
			},
		).
		Returning("id").
		ToSQL()
}
