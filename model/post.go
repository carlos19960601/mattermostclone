package model

import (
	"encoding/json"
	"fmt"
	"io"
)

type Post struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	ChannelId string `json:"channel_id"`

	Message string `json:"message"`
}

func (p *Post) Clone() *Post {
	copy := Post{}
	_ = p.ShallowCopy(&copy)
	return &copy
}

func (p *Post) ToJSON() string {
	copy := p.Clone()
	b, _ := json.Marshal(copy)
	return string(b)
}

func (p *Post) ShallowCopy(dst *Post) error {
	if dst == nil {
		return fmt.Errorf("目标不能为nil")
	}
	dst.Id = p.Id
	return nil
}

func PostFromJson(data io.Reader) *Post {
	var post Post
	_ = json.NewDecoder(data).Decode(&post)
	return &post
}
