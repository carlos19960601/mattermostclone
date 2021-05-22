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
	Roles    string `json:"roles"`
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

func (u *User) GetDisplayNameWithPrefix(nameFormat, prefix string) string {
	displayName := prefix + u.Username
	return u.getDisplayName(displayName, nameFormat)
}

func (u *User) getDisplayName(baseName, nameFormat string) string {
	displayName := baseName
	return displayName
}

func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func (u *User) PreSave() {
	if u.Id == "" {
		u.Id = NewId()
	}
	if u.Username == "" {
		u.Username = NewId()
	}

	if u.Password != "" {
		u.Password = HashPassword(u.Password)
	}
}
