package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tjerman/lq/main/game"
)

func main() {
	c := game.GameConfig{
		Title:    "LQ: The Game",
		TileSize: 40,
	}

	g := game.Init(c)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
