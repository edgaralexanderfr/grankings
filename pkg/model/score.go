package model

import (
	"github.com/edgaralexanderfr/grankings/pkg/storage"
)

type Score struct {
	Score  int `json:"score"`
	Player `json:"player"`
}

func (score *Score) Create() {
	storage.Insert("scores", score)
}
