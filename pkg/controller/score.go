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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(scores)
}

func ScoreCreate(w http.ResponseWriter, r *http.Request) {
	var score model.Score = model.Score{
		Score: 10000,
		Player: model.Player{
			Name:         "Edgar Alexander Franco",
			Abbreviation: "EAF",
		},
	}

	score.Create()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(score)
}
