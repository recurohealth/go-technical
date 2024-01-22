package model

import "time"

type Address struct {
	Id             string     `json:"id" db:"id"`
	City           string     `json:"city,omitempty" db:"city"`
	AddressLineOne string     `json:"addressLineOne,omitempty" db:"address_line_one"`
	AddressLineTwo *string    `json:"addressLineTwo,omitempty" db:"address_line_two"`
	State          string     `json:"state,omitempty" db:"state"`
	Zip            string     `json:"zip,omitempty" db:"zip"`
	Country        string     `json:"country,omitempty" db:"country"`
	DateCreated    *time.Time `json:"dateCreated,omitempty" db:"date_created"`
	DateModified   *time.Time `json:"dateModified,omitempty" db:"date_modified"`
	PatientId      string     `json:"patientId" db:"patient_id"`
}
