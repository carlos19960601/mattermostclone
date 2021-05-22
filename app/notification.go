package app

import "github.com/zengqiang96/mattermostclone/model"

type PostNotification struct {
	Channel *model.Channel
	Post    *model.Post
	Sender  *model.User
}

func (a *App) SendNotification(post *model.Post, team *model.Team, channel *model.Channel, sender *model.User) error {
	if channel.DeleteAt > 0 {
		return nil
	}

	if channel.Type == model.CHANNEL_DIRECT {
		_ = channel.GetOtherUserIdFromDM(post.UserId)
	}

	notification := &PostNotification{
		Post:    post.Clone(),
		Channel: channel,
		Sender:  sender,
	}

	message := model.NewWebSocketEvent(model.WEBSOCKET_EVENT_POSTED, "", post.ChannelId, "")
	message.Add("post", post.ToJSON())
	message.Add("channel_type", channel.Type)
	message.Add("channel_display_name", notification.GetChannelName(model.SHOW_USERNAME))
	message.Add("channel_name", channel.Name)
	message.Add("sender_name", notification.GetSenderName(model.SHOW_USERNAME))
	message.Add("team_id", team.Id)

	a.Publish(message)
	return nil
}

func (n *PostNotification) GetSenderName(usernameFormat string) string {
	return n.Sender.GetDisplayNameWithPrefix(usernameFormat, "@")
}

func (n *PostNotification) GetChannelName(usernameFormat string) string {
	switch n.Channel.Type {
	case model.CHANNEL_DIRECT:
		return n.Sender.GetDisplayNameWithPrefix(usernameFormat, "@")
	default:
		return n.Channel.DisplayName
	}

}
