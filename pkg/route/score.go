package route

import (
	"net/http"

	"github.com/edgaralexanderfr/grankings/pkg/controller"
	"github.com/gorilla/mux"
)

func ScoreCreateRoute(r *mux.Router) {
	r.
		Name("ScoreList").
		Methods("GET").
		Path("/scores").
		Handler(http.HandlerFunc(controller.ScoreList))

	r.
		Name("ScoreCreate").
		Methods("POST").
		Path("/scores").
		Handler(http.HandlerFunc(controller.ScoreCreate))
}
