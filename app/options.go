package app

import "github.com/zengqiang96/mattermostclone/config"

type Option func(s *Server) error

func ConfigStore(configStore *config.Store) Option {
	return func(s *Server) error {
		s.configStore = configStore
		return nil
	}
}

type AppOption func(s *App)

func ServerConnector(s *Server) AppOption {
	return func(a *App) {
		a.srv = s
	}
}
