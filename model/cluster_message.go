package model

import (
	"encoding/json"
	"io"
)

const (
	CLUSTER_EVENT_PUBLISH = "publish"
)

type ClusterMessage struct {
	Event    string `json:"event"`
	SendType string `json:"-"`
	Data     string `json:"data,omitempty"`
}

func (o *ClusterMessage) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func ClusterMessageFromJSON(data io.Reader) *ClusterMessage {
	var o ClusterMessage
	_ = json.NewDecoder(data).Decode(&o)
	return &o
}
