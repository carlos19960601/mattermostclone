package api4

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/model"
)

func (api *API) InitUser() {
	api.BaseRoutes.Users.Handle("", api.ApiHandler(createUser)).Methods("POST")
	api.BaseRoutes.Users.Handle("", api.ApiSessionRequired(getUsers)).Methods("GET")

	api.BaseRoutes.Users.Handle("/login", api.ApiHandler(login)).Methods("POST")
}

func createUser(ctx *Context, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJson(r.Body)
	if user == nil {
		ctx.SetInvalidParam("user")
		return
	}

	ruser, err := ctx.App.CreateUserFromSignup(ctx.AppContext, user)
	if err != nil {
		ctx.Err = err
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(ruser.ToJSON()))
}

func getUsers(ctx *Context, w http.ResponseWriter, r *http.Request) {

}

func login(ctx *Context, w http.ResponseWriter, r *http.Request) {
	props := model.MapFromJSON(r.Body)
	loginId := props["login_id"]
	password := props["password"]

	user, err := ctx.App.AuthenticateUserForLogin(ctx.AppContext, loginId, password)
	if err != nil {
		ctx.Err = err
		return
	}

	err = ctx.App.DoLogin(ctx.AppContext, w, r, user)
	if err != nil {
		ctx.Err = err
		return
	}
	_, _ = w.Write([]byte(user.ToJSON()))
}
