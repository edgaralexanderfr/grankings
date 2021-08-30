package model

import (
	"github.com/edgaralexanderfr/grankings/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapScores(data []storage.D) (scores []Score) {
	scores = []Score{}

	for _, row := range data {
		scores = append(scores, MapScore(row))
	}

	return
}

func MapScore(data storage.D) (score Score) {
	score = Score{}

	for _, col := range data {
		switch col.Key {
		case "_id":
			score.Id = col.Value.(primitive.ObjectID).Hex()
		case "score":
			score.Score = int(col.Value.(int32))
		case "player":
			score.Player = MapPlayer(storage.PrimitiveDToStorageD(col.Value.(primitive.D)))
		}
	}

	return
}
