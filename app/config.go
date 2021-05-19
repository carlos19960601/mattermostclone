package app

import "github.com/zengqiang96/mattermostclone/model"

func (s *Server) Config() *model.Config {
	return s.configStore.Get()
}
