package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type (
	GameConfig struct {
		Title string

		TileSize uint
		SizeX    int
		SizeY    int
		// @todo...
	}

	Game struct {
		config GameConfig
		// @todo...
	}
)

// Init inits the Game struct
func Init(c GameConfig) *Game {
	title := c.Title
	ebiten.SetWindowTitle(title)

	sx := c.SizeX
	sy := c.SizeY
	if sx <= 0 {
		sx = 640
	}
	if sy <= 0 {
		sy = 480
	}

	ebiten.SetWindowSize(sx, sy)

	return &Game{
		config: c,
	}
}

// Update handles the game logic update cycle
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw handles the game draw cycle
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout defines the game's logical size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
