package model

import (
	"encoding/json"
	"io"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
	DeleteAt int64  `json:"delete_at"`
}

func UserFromJson(data io.Reader) *User {
	var user User
	_ = json.NewDecoder(data).Decode(&user)
	return &user
}

func (u *User) ToJSON() string {
	b, _ := json.Marshal(u)
	return string(b)
}

func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
