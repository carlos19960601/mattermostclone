package web

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/app"
	"github.com/zengqiang96/mattermostclone/app/request"
)

type Handler struct {
	App        app.AppIface
	HandleFunc func(*Context, http.ResponseWriter, *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		App:        h.App,
		AppContext: &request.Context{},
	}

	h.HandleFunc(c, w, r)
}
