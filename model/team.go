package model

type Team struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
	DeleteAt    int64  `json:"delete_at"`
}
