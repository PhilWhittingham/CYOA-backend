package main

import (
	"github.com/PhilWhittingham/CYOA-backend/api"
	"github.com/PhilWhittingham/CYOA-backend/player"
)

func main() {
	pID := "1"
	pd := player.NewPlayer(pID)
	api.Initialise(pd)
}
