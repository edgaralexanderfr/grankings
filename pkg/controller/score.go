package controller

import (
	"net/http"
	"os"
	"strconv"

	ghttp "github.com/edgaralexanderfr/grankings/pkg/http"
	"github.com/edgaralexanderfr/grankings/pkg/model"
)

func ScoreList(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(os.Getenv("GET_LIST_DEFAULT_LIMIT"))

	if err != nil {
		limit = 30
	}

	list, _ := model.FindScores(0, limit)

	ghttp.Response(w, r, list, http.StatusOK)
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

	ghttp.Response(w, r, score, http.StatusCreated)
}
