package app

import (
	"github.com/gorilla/websocket"
	"github.com/zengqiang96/mattermostclone/model"
)

type WebConnConfig struct {
	Websocket *websocket.Conn
	Session   model.Session
}

type WebConn struct {
	App       *App
	Websocket *websocket.Conn
	UserId    string
}

func (wc *WebConn) Close() {
	wc.Websocket.Close()
}

func (a *App) NewWebConn(cfg *WebConnConfig) *WebConn {
	if cfg.Session.UserId == "" {

	}

	wc := WebConn{
		App:       a,
		Websocket: cfg.Websocket,
		UserId:    cfg.Session.UserId,
	}
	return &wc
}
