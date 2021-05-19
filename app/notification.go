package app

import "github.com/zengqiang96/mattermostclone/model"

func (a *App) SendNotification(post *model.Post, team *model.Team, channel *model.Channel, sender *model.User) error {
	if channel.DeleteAt > 0 {
		return nil
	}

	if channel.Type == model.CHANNEL_DIRECT {
		_ = channel.GetOtherUserIdFromDM(post.UserId)
	}

	message := model.NewWebSocketEvent(model.WEBSOCKET_EVENT_POSTED, "", post.ChannelId, "")
	message.Add("post", post.ToJSON())
	message.Add("channel_type", channel.Type)
	message.Add("channel_name", channel.Name)

	a.Publish(message)
	return nil
}
