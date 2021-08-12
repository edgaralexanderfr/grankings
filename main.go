package main

import (
	"fmt"

	"github.com/edgaralexanderfr/grankings/pkg/model"
)

func main() {
	player := model.Player{
		Name:         "Edgar Alexander Franco",
		Abbreviation: "EAF",
		Key:          "SECRET_KEY",
	}

	fmt.Println("Grankings v0.0.1")
	fmt.Println("")
	fmt.Println("Top players:")
	fmt.Println("1. " + player.GetAbbreviation())
}
