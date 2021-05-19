package model

import (
	"encoding/json"
	"io"
)

type AppError struct {
	Id         string `json:"id"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewAppError(where string, id string, param map[string]interface{}, details string, status int) *AppError {
	ap := AppError{}
	ap.Id = id
	ap.Message = id
	ap.StatusCode = status
	return &ap
}

func MapFromJSON(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	}
	return objmap
}
