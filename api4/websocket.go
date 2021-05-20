package api4

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zengqiang96/mattermostclone/app"
	"github.com/zengqiang96/mattermostclone/model"
)

func (api *API) InitWebsocket() {
	api.BaseRoutes.ApiRoot.Handle("/{websocket:websocket(?:\\/)?}", api.ApiHanderTrustRequest(connectWebsocket)).Methods("GET")
}

func connectWebsocket(c *Context, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  model.SOCKET_MAX_MESSAGE_SIZE_KB,
		WriteBufferSize: model.SOCKET_MAX_MESSAGE_SIZE_KB,
		CheckOrigin:     c.App.OriginChecker(),
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.Err = model.NewAppError("connect", "api.websocket.connect.upgrade.app_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	cfg := &app.WebConnConfig{
		Websocket: ws,
		Session:   *c.AppContext.Session(),
	}

	wc := c.App.NewWebConn(cfg)
	if c.AppContext.Session().UserId != "" {
		c.App.HubRegister(wc)
	}
}
