package database

import (
	"github.com/doug-martin/goqu/v9"
)

const ADDRESS_TABLE = "address"

type AddressQuery struct{}

func (AddressQuery) InsertAddress() (query string, args []interface{}, err error) {
	return postgresDialect.Insert(ADDRESS_TABLE).
		Prepared(true).
		Rows(
			goqu.Record{
				"id":               goqu.L("$1"),
				"city":             goqu.L("$2"),
				"address_line_one": goqu.L("$3"),
				"address_line_two": goqu.L("$4"),
				"state":            goqu.L("$5"),
				"zip":              goqu.L("$6"),
				"country":          goqu.L("$7"),
				"patient_id":       goqu.L("$8"),
			},
		).
		Returning("id").
		ToSQL()
}
