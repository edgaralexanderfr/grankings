package model

import (
	"github.com/edgaralexanderfr/grankings/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapPlayers(data []storage.D) (players []Player) {
	players = []Player{}

	for _, row := range data {
		players = append(players, MapPlayer(row))
	}

	return
}

func MapPlayer(data storage.D) (player Player) {
	player = Player{}

	for _, col := range data {
		switch col.Key {
		case "_id":
			player.Id = col.Value.(primitive.ObjectID).Hex()
		case "name":
			player.Name = col.Value.(string)
		case "abbreviation":
			player.Abbreviation = col.Value.(string)
		case "key":
			player.Key = col.Value.(string)
		}
	}

	return
}
