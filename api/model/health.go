package model

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type HealthMessage struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewHealthMessage(statusCode int, message string) HealthMessage {
	healthMessage := HealthMessage{
		Id:      uuid.New().String(),
		Status:  fmt.Sprintf("%v %v", statusCode, http.StatusText(statusCode)),
		Message: message,
	}

	return healthMessage
}
