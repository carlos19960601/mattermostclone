package app

import (
	"errors"
	"fmt"
	"hash/maphash"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/zengqiang96/mattermostclone/config"
	"github.com/zengqiang96/mattermostclone/einterfaces"
	"github.com/zengqiang96/mattermostclone/services/cache"
	"github.com/zengqiang96/mattermostclone/store"
	"github.com/zengqiang96/mattermostclone/store/localcachelayer"
	"github.com/zengqiang96/mattermostclone/store/retrylayer"
	"github.com/zengqiang96/mattermostclone/store/searchlayer"
	"github.com/zengqiang96/mattermostclone/store/sqlstore"
	"github.com/zengqiang96/mattermostclone/store/timerlayer"
)

type Server struct {
	sqlStore *sqlstore.SqlStore
	Store    store.Store

	RootRouter *mux.Router
	Router     *mux.Router

	Server     *http.Server
	ListenAddr *net.TCPAddr

	hubs     []*Hub
	hashSeed maphash.Seed

	newStore func() (store.Store, error)

	configStore *config.Store

	sessionCache cache.Cache

	Cluster einterfaces.ClusterInterface
}

func NewServer(options ...Option) (*Server, error) {
	rootRouter := mux.NewRouter()
	s := Server{
		RootRouter: rootRouter,
	}

	for _, option := range options {
		if err := option(&s); err != nil {
			return nil, fmt.Errorf("option启用失败  err: %w", err)
		}
	}

	fakeApp := New(ServerConnector(&s))
	fakeApp.HubStart()

	var err error

	if s.newStore == nil {
		s.newStore = func() (store.Store, error) {
			s.sqlStore = sqlstore.New(s.Config().SqlSettings)

			lcl, err := localcachelayer.NewLocalCacheLayer(retrylayer.New(s.sqlStore))
			if err != nil {
				return nil, fmt.Errorf("创建localcachelayer失败 err: %w", err)
			}

			searchStore := searchlayer.NewSearchLayer(lcl)

			return timerlayer.New(searchStore), nil
		}
	}

	s.Store, err = s.newStore()
	if err != nil {
		return nil, fmt.Errorf("创建store失败, err: %w", err)
	}

	s.Router = s.RootRouter.PathPrefix("/").Subrouter()
	return &s, nil
}

func (s *Server) Start() error {
	var handler http.Handler = s.RootRouter
	s.Server = &http.Server{
		Handler: handler,
	}

	addr := s.Config().ServiceSettings.ListenAddress
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("服务启动失败 error: %w", err)
	}
	s.ListenAddr = listener.Addr().(*net.TCPAddr)

	go func() {
		err := s.Server.Serve(listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("启动服务失败 err: %v\n", err)
			time.Sleep(time.Second)
		}
	}()

	return nil
}

func (s *Server) Shutdown() {
	s.configStore.Close()
}

func (s *App) OriginChecker() func(r *http.Request) bool {
	return nil
}
