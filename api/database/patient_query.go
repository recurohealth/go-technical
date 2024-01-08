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
				"member_id":     goqu.L("$2"),
				"client_id":     goqu.L("$3"),
				"first_name":    goqu.L("$4"),
				"last_name":     goqu.L("$5"),
				"gender":        goqu.L("$6"),
				"date_of_birth": goqu.L("$7"),
				"phone_number":  goqu.L("$8"),
				"email":         goqu.L("$9"),
			},
		).
		Returning("id").
		ToSQL()
}
