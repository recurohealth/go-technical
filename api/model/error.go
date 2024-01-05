package model

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type ErrorMessage struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewErrorMessage(statusCode int, message string, err error) ErrorMessage {
	errorMessage := ErrorMessage{
		Id:      uuid.New().String(),
		Status:  fmt.Sprintf("%v %v", statusCode, http.StatusText(statusCode)),
		Message: message,
		Error:   err.Error(),
	}

	return errorMessage
}
