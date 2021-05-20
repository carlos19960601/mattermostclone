package api4

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/web"
)

type Context = web.Context

func (api *API) ApiHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	handler := web.Handler{
		App:        api.app,
		HandleFunc: h,
	}
	return handler
}

func (api *API) ApiSessionRequired(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	handler := web.Handler{
		App:        api.app,
		HandleFunc: h,
	}
	return &handler
}

func (api *API) ApiHanderTrustRequest(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	handler := web.Handler{
		App:        api.app,
		HandleFunc: h,
	}
	return &handler
}
