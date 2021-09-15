package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/model"
)

func (a *App) CreatePostAsUser(c *request.Context, post *model.Post, currentSessionId string) (*model.Post, *model.AppError) {
	channel, errCh := a.Srv().Store.Channel().Get(post.ChannelId)
	if errCh != nil {
		err := model.NewAppError("CreatePostAsUser", "api.context.invalid_param.app_error", map[string]interface{}{"Name": "post.channel_id"}, errCh.Error(), http.StatusBadRequest)
		return nil, err
	}
	rp, err := a.CreatePost(c, post, channel)
	if err != nil {
		// TODO
	}
	return rp, nil
}

func (a *App) CreatePost(c *request.Context, post *model.Post, channel *model.Channel) (savedPost *model.Post, err *model.AppError) {
	user, nErr := a.Srv().Store.User().Get(context.Background(), post.UserId)
	if nErr != nil {
		// TODO:
	}

	rpost, nErr := a.Srv().Store.Post().Save(post)
	if nErr != nil {
		// TODO:
	}

	if err := a.handlePostEvents(c, rpost, user, channel); err != nil {
		fmt.Println("处理Post Events失败 err: %w", err)
	}
	return rpost, nil
}

func (a *App) handlePostEvents(c *request.Context, post *model.Post, user *model.User, channel *model.Channel) error {
	var team *model.Team
	if channel.TeamId != "" {
		t, err := a.Srv().Store.Team().Get(channel.TeamId)
		if err != nil {
			return err
		}
		team = t
	} else {
		team = &model.Team{}
	}
	if err := a.SendNotification(post, team, channel, user); err != nil {
		return err
	}
	return nil
}
