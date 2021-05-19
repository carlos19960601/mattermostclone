package app

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/zengqiang96/mattermostclone/config"
	"github.com/zengqiang96/mattermostclone/store"
)

type Server struct {
	Store store.Store

	RootRouter *mux.Router
	Router     *mux.Router

	Server     *http.Server
	ListenAddr *net.TCPAddr

	hubs []*Hub

	configStore *config.Store
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
