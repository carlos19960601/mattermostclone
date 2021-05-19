package app

import "github.com/zengqiang96/mattermostclone/model"

func (a *App) CreateSession(session *model.Session) (*model.Session, *model.AppError) {
	session.Token = ""

	session, err := a.Srv().Store.Session().Save(session)
	if err != nil {
		
	}
	return session, nil
}
