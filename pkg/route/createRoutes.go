package route

import (
	"github.com/gorilla/mux"
)

func CreateRoutes() (r *mux.Router) {
	r = mux.NewRouter().StrictSlash(true)

	ScoreCreateRoute(r)

	return
}
