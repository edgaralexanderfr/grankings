package router

import (
	"github.com/edgaralexanderfr/grankings/pkg/route"
	"github.com/gorilla/mux"
)

func CreateRoutes() (r *mux.Router) {
	r = mux.NewRouter().StrictSlash(true)

	route.ScoreCreateRoute(r)

	return
}
