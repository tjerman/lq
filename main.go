package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/tjerman/lq/game"
	"github.com/tjerman/lq/level"
)

func main() {
	setup()

	c := &game.GameConfig{
		Title:    "lq",
		TileSize: 40,
		Levels:   []game.Level{level.Load()},
	}

	env(c)

	g, err := game.Init(c)
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// Setup should perform any system initialization that we may need
func setup() {
	// Bits regarding random values
	// rand.Seed(time.Now().UnixNano())
}

func env(c *game.GameConfig) {
	if os.Getenv("DEBUG_GRID") == "1" {
		c.DebugGrid = true
	}

	if os.Getenv("DEBUG_OBJECT_LABELS") == "1" {
		c.DebugObjectLabels = true
	}

	if os.Getenv("DEBUG_ROOM_STRUCTURE") == "1" {
		c.DebugRoomStructure = true
	}

	if os.Getenv("DEBUG_CORRIDOR_STRUCTURE") == "1" {
		c.DebugCorridorStructure = true
	}

	if os.Getenv("DEBUG_FIXED_WORLD") == "1" {
		c.DebugFixedWorld = true
	}

	if os.Getenv("DEBUG_INPUT_CONTROLLS") == "1" {
		c.DebugInputControlls = true
	}
}
