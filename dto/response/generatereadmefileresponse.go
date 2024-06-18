package dto

import (
	"time"
)

type GenereteReadmeFileResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
