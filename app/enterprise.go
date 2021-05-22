package app

import "github.com/zengqiang96/mattermostclone/einterfaces"

var clusterInterface func(*Server) einterfaces.ClusterInterface

func (s *Server) initEnterprise() {
	if clusterInterface != nil {
		s.Cluster = clusterInterface(s)
	}
}
