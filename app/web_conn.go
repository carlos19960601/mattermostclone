package app

import "github.com/gorilla/websocket"

type WebConn struct {
	App       *App
	WebSocket *websocket.Conn
	UserId    string
}

func (wc *WebConn) Close() {
	wc.WebSocket.Close()
}
