package model

import "strings"

type Player struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Key          string `json:"-"`
}

func (player *Player) GetAbbreviation() string {
	parts := strings.Split(player.Name, " ")

	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(string(part[0]))
		} else {
			parts[i] = ""
		}
	}

	return strings.Join(parts, "")
}
