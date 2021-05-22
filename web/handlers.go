package web

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/app"
	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/utils"
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

	token, tokenLocation := app.ParseAuthTokenFromRequest(r)
	if token != "" && tokenLocation != app.TokenLocationCloudHeader {
		session, err := c.App.GetSession(token)
		defer app.ReturnSessionToPool(session)

		if err != nil {

		}
		c.AppContext.SetSession(session)
	}

	if c.Err == nil {
		h.HandleFunc(c, w, r)
	}

	if c.Err != nil {
		utils.RenderWebAppError(c.App.Config(), w, r, c.Err)
	}
}
