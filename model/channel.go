package model

import "strings"

const (
	CHANNEL_DIRECT = "D"
)

type Channel struct {
	Id          string `json:"id"`
	TeamId      string `json:"team_id"`
	Type        string `json:"type"`
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	Header      string `json:"header"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
	DeleteAt    int64  `json:"delete_at"`
}

// 2个人之间的聊天 channel.name = userId1__userId2
func (c *Channel) GetOtherUserIdFromDM(userId string) string {
	if c.Type != CHANNEL_DIRECT {
		return ""
	}
	userIds := strings.Split(c.Name, "__")

	var otherUserId string
	if userIds[0] != userIds[1] {
		if userIds[0] == userId {
			otherUserId = userIds[1]
		} else {
			otherUserId = userIds[0]
		}
	}
	return otherUserId
}
