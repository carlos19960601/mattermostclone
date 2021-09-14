package sqlstore

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zengqiang96/mattermostclone/model"
	"github.com/zengqiang96/mattermostclone/store"
)

type SqlSessionStore struct {
	*SqlStore
}

func (me SqlSessionStore) Save(session *model.Session) (*model.Session, error) {
	if session.Id != "" {
		return nil, store.NewErrInvalidInput("Session", "id", session.Id)
	}

	session.PreSave()

	if err := me.GetMaster().Insert(session); err != nil {
		return nil, errors.Wrapf(err, "failed to save Session with id=%s", session.Id)
	}

	teamMembers, err := me.Team().GetTeamsForUser(context.Background(), session.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find TeamMembers for Session with userId=%s", session.UserId)
	}

	session.TeamMembers = make([]*model.TeamMember, 0, len(teamMembers))
	for _, tm := range teamMembers {
		if tm.DeleteAt == 0 {
			session.TeamMembers = append(session.TeamMembers, tm)
		}
	}

	return session, nil
}
