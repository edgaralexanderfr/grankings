package model

import (
	"github.com/edgaralexanderfr/grankings/pkg/storage"
)

type Score struct {
	Id     string `json:"id"`
	Score  int    `json:"score"`
	Player `json:"player"`
}

func (score *Score) Create() {
	storage.Insert("scores", score)
}

func FindScores(start int, limit int) (scores []Score, err error) {
	results, err := storage.Find("scores", start, limit)

	if err != nil {
		return
	}

	scores = MapScores(results)

	return
}
