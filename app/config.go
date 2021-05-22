package app

import "github.com/zengqiang96/mattermostclone/model"

func (s *Server) Config() *model.Config {
	return s.configStore.Get()
}

func (a *App) Config() *model.Config {
	return a.srv.Config()
}

func (a *App) GetCookieDomain() string {
	return ""
}
