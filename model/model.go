package model

import (
	"event-management/handler/errors"
	"time"
)

type Timestamps struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Response struct {
	Message string `json:"message"`
}

type ResponseWithData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWithError struct {
	Message string                  `json:"message"`
	Error   errors.ErrorValidations `json:"error"`
}
