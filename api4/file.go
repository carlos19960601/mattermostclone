package api4

func (api *API) InitFile() {
	api.BaseRoutes.Files.Handle("", api.ApiSessionRequired(createTeam)).Methods("POST")
	api.BaseRoutes.Files.Handle("", api.ApiSessionRequired(getAllTeams)).Methods("GET")
}
