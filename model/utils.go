package model

import (
	"bytes"
	"encoding/base32"
	"encoding/json"
	"io"
	"time"

	"github.com/pborman/uuid"
)

type AppError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`
	DetailedError string `json:"detailed_error"`
	StatusCode    int    `json:"status_code"`
	Where         string `json:"-"`
}

func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}

func NewAppError(where string, id string, param map[string]interface{}, details string, status int) *AppError {
	ap := AppError{}
	ap.Id = id
	ap.Message = id
	ap.StatusCode = status
	ap.DetailedError = details
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

func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	b.Truncate(26) // removes "=="
	return b.String()
}
