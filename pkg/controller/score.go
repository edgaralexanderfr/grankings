package controller

import (
	"encoding/json"
	"net/http"

	"github.com/edgaralexanderfr/grankings/pkg/model"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(scores)
}
