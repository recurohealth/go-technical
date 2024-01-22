package model

const (
	MALE    Gender = "M"
	FEMALE  Gender = "F"
	UNKNOWN Gender = "U"
)

type Gender string

type ClientPatientInfo struct {
	ClientId    string  `json:"clientId" db:"client_id"`
	FirstName   string  `json:"firstName" db:"first_name"`
	LastName    string  `json:"lastName" db:"last_name"`
	Gender      Gender  `json:"gender" db:"gender"`
	DateOfBirth string  `json:"dateOfBirth" db:"date_of_birth"`
	PhoneNumber string  `json:"phoneNumber" db:"phone_number"`
	Email       string  `json:"email" db:"email"`
	Address     Address `json:"address" db:"address"`
}

type Patient struct {
	Id           string `json:"id" db:"id"`
	ClientId     string `json:"clientId,omitempty" db:"client_id"`
	FirstName    string `json:"firstName,omitempty" db:"first_name"`
	LastName     string `json:"lastName,omitempty" db:"last_name"`
	Gender       Gender `json:"gender,omitempty" db:"gender"`
	DateOfBirth  string `json:"dateOfBirth,omitempty" db:"date_of_birth"`
	PhoneNumber  string `json:"phoneNumber,omitempty" db:"phone_number"`
	Email        string `json:"email,omitempty" db:"email"`
	DateCreated  string `json:"dateCreated,omitempty" db:"date_created"`
	DateModified string `json:"dateModified,omitempty" db:"date_modified"`
}
