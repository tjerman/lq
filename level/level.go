package level

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/tjerman/lq/game"
)

type (
	Level struct {
		Rooms     []*Room
		Corridors []*Corridor
		config    *game.GameConfig
	}
)

func Load() *Level {
	return &Level{}
}

func (l *Level) String() string {
	return "level"
}

func (l *Level) Init(g *game.Game) error {
	l.config = g.Config

	var rooms []*Room
	var err error
	if l.config.DebugFixedWorld {
		rooms, err = determineRoomsF(l.config)
	} else {
		// @todo...
	}
	if err != nil {
		return err
	}
	l.Rooms = rooms

	corridors, err := determineCorridors(l.config, nil, rooms)
	if err != nil {
		return err
	}
	l.Corridors = corridors

	return nil
}

func (l *Level) Update(screen *ebiten.Image) error {
	return nil
}

func (l *Level) Draw(screen *ebiten.Image) {
	if l.config.DebugRoomStructure {
		drawRooms(screen, l.Rooms)
	}

	if l.config.DebugCorridorStructure {
		drawCorridors(screen, l.Corridors)
	}
}
