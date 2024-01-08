package model

const (
	MALE    Gender = "M"
	FEMALE  Gender = "F"
	UNKNOWN Gender = "U"
)

type Gender string

type Patient struct {
	Id           string `json:"id" db:"id"`
	ClientId     string `json:"clientId" db:"client_id"`
	FirstName    string `json:"firstName" db:"first_name"`
	LastName     string `json:"lastName" db:"last_name"`
	Gender       Gender `json:"gender" db:"gender"`
	DateOfBirth  string `json:"dateOfBirth" db:"date_of_birth"`
	PhoneNumber  string `json:"phoneNumber" db:"phone_number"`
	Email        string `json:"email" db:"email"`
	DateCreated  string `json:"dateCreated" db:"date_created"`
	DateModified string `json:"dateModified" db:"date_modified"`
}
