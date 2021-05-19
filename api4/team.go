package api4

import "net/http"

func (api *API) InitTeam() {
	api.BaseRoutes.Teams.Handle("", api.ApiSessionRequired(createTeam)).Methods("POST")
	api.BaseRoutes.Teams.Handle("", api.ApiSessionRequired(getAllTeams)).Methods("GET")
}

func createTeam(ctx *Context, w http.ResponseWriter, r *http.Request) {
	
}

func getAllTeams(ctx *Context, w http.ResponseWriter, r *http.Request) {

}
