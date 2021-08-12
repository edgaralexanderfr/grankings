package model

import "strings"

type Player struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Key          string `json:"key"`
}

func (player *Player) GetAbbreviation() string {
	parts := strings.Split(player.Name, " ")

	for i, part := range parts {
		parts[i] = strings.ToUpper(string(part[0]))
	}

	return strings.Join(parts, "")
}
