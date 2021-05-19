package model

const (
	WEBSOCKET_EVENT_POSTED = "posted"
)

type WebSocketEvent struct {
	event     string
	data      map[string]interface{}
	Broadcast *WebSocketBroadcast
}

type WebSocketBroadcast struct {
	UserId    string `json:"user_id"`
	TeamId    string `json:"team_id"`
	ChannelId string `json:"channel_id"`
}

func NewWebSocketEvent(event, teamId, channelId, userId string) *WebSocketEvent {
	return &WebSocketEvent{event: event, data: make(map[string]interface{}),
		Broadcast: &WebSocketBroadcast{UserId: userId, TeamId: teamId, ChannelId: channelId}}
}

func (ev *WebSocketEvent) Add(key string, value interface{}) {
	ev.data[key] = value
}

func (ev *WebSocketEvent) GetBroadcast() *WebSocketBroadcast {
	return ev.Broadcast
}
