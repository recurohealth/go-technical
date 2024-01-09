package model

import "time"

type Address struct {
	Id             string     `json:"id" db:"id"`
	City           string     `json:"city" db:"city"`
	AddressLineOne string     `json:"addressLineOne" db:"address_line_one"`
	AddressLineTwo *string    `json:"addressLineTwo" db:"address_line_two"`
	State          string     `json:"state" db:"state"`
	Zip            string     `json:"zip" db:"zip"`
	Country        string     `json:"country" db:"country"`
	DateCreated    *time.Time `json:"dateCreated" db:"date_created"`
	DateModified   *time.Time `json:"dateModified" db:"date_modified"`
	PatientId      string     `json:"patientId" db:"patient_id"`
}
