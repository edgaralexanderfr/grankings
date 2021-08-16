package route

import (
	"encoding/json"
	"net/http"

	"github.com/edgaralexanderfr/grankings/pkg/model"
	"github.com/gorilla/mux"
)

func ScoreList(w http.ResponseWriter, r *http.Request) {
	scores := []model.Score{
		{
			Score: 10000,
			Player: model.Player{
				Name:         "Edgar Alexander Franco",
				Abbreviation: "EAF",
			},
		},
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(scores)
}

func ScoreCreateRoute(r *mux.Router) {
	r.
		Name("ScoreList").
		Methods("GET").
		Path("/scores").
		Handler(http.HandlerFunc(ScoreList))
}
