package api4

import (
	"github.com/gorilla/mux"
	"github.com/zengqiang96/mattermostclone/app"
	"github.com/zengqiang96/mattermostclone/model"
)

type Routes struct {
	Root    *mux.Router // ''
	ApiRoot *mux.Router // 'api/v4'

	Users *mux.Router // 'api/v4/users'
	User  *mux.Router // 'api/v4/users/{user_id:[A-Za-z0-9]+}'

	Teams *mux.Router // 'api/v4/teams'
	Team  *mux.Router // 'api/v4/teams/{team_id:[A-Za-z0-9]+}'

	Posts *mux.Router // 'api/v4/posts'
	Post  *mux.Router // 'api/v4/posts/{post_id:[A-Za-z0-9]+}'

	Files *mux.Router // 'api/v4/files'
	File  *mux.Router // 'api/v4/files/{file_id:[A-Za-z0-9]+}'
}

type API struct {
	app        app.AppIface
	BaseRoutes *Routes
}

func Init(a app.AppIface, root *mux.Router) *API {
	api := API{
		app:        a,
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.Root = root
	api.BaseRoutes.ApiRoot = root.PathPrefix(model.API_URL_SUFFIX_V4).Subrouter()

	api.BaseRoutes.Users = api.BaseRoutes.ApiRoot.PathPrefix("/users").Subrouter()
	api.BaseRoutes.User = api.BaseRoutes.Users.PathPrefix("/{user_id:[A-Za-z0-9]+}").Subrouter()

	api.BaseRoutes.Teams = api.BaseRoutes.ApiRoot.PathPrefix("/teams").Subrouter()
	api.BaseRoutes.Team = api.BaseRoutes.Teams.PathPrefix("/{team_id:[A-Za-z0-9]+}").Subrouter()

	api.BaseRoutes.Posts = api.BaseRoutes.ApiRoot.PathPrefix("/posts").Subrouter()
	api.BaseRoutes.Team = api.BaseRoutes.Posts.PathPrefix("/{post_id:[A-Za-z0-9]+}").Subrouter()

	api.BaseRoutes.Files = api.BaseRoutes.ApiRoot.PathPrefix("/files").Subrouter()
	api.BaseRoutes.File = api.BaseRoutes.Files.PathPrefix("/{file_id:[A-Za-z0-9]+}").Subrouter()

	api.InitUser()
	api.InitTeam()
	api.InitPost()
	api.InitFile()

	return &api
}
