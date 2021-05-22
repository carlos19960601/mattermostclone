package einterfaces

import "github.com/zengqiang96/mattermostclone/model"

type ClusterInterface interface {
	SendClusterMessage(msg *model.ClusterMessage)
}


