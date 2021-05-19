package api4

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/model"
)

func (api *API) InitPost() {
	api.BaseRoutes.Posts.Handle("", api.ApiSessionRequired(createPost)).Methods("POST")
	api.BaseRoutes.Posts.Handle("", api.ApiSessionRequired(getPost)).Methods("GET")
}

func createPost(ctx *Context, w http.ResponseWriter, r *http.Request) {
	post := model.PostFromJson(r.Body)
	if post == nil {
		ctx.SetInvalidParam("post")
		return
	}

	post.UserId = ctx.AppContext.Session().UserId
	rp, err := ctx.App.CreatePostAsUser(ctx.AppContext, post, ctx.AppContext.Session().Id)
	if err != nil {
		ctx.Err = err
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(rp.ToJSON()))
}

func getPost(ctx *Context, w http.ResponseWriter, r *http.Request) {

}
