package game

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type (
	GameConfig struct {
		Title string

		// Size in pixles on both axises
		SizeX int
		SizeY int

		// Grid length on both axises
		TileSize int
		LenX     int
		LenY     int

		Levels []Level

		DebugGrid              bool
		DebugObjectLabels      bool
		DebugRoomStructure     bool
		DebugCorridorStructure bool
		DebugFixedWorld        bool
	}

	Game struct {
		Config *GameConfig

		Level Level
	}

	Level interface {
		Init(g *Game) error
		Update(screen *ebiten.Image) error
		Draw(screen *ebiten.Image)
	}
)

var (
	ErrorNoLevelsDefined = errors.New("game: no levels defined")
)

// Init inits the Game struct
func Init(c *GameConfig) (*Game, error) {
	// Some validity checks
	if c.Levels == nil {
		return nil, ErrorNoLevelsDefined
	}

	// Some meta
	sx := c.SizeX
	sy := c.SizeY
	if sx <= 0 {
		sx = 640
		c.SizeX = sx
	}
	if sy <= 0 {
		sy = 480
		c.SizeY = sy
	}

	c.LenX = int(c.SizeX / c.TileSize)
	c.LenY = int(c.SizeY / c.TileSize)

	title := c.Title
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(sx, sy)

	// @todo this should be something nicer
	g := &Game{Config: c, Level: c.Levels[0]}

	for _, l := range c.Levels {
		l.Init(g)
	}

	return g, nil
}

func (g *Game) String() string {
	return fmt.Sprintf("game: %s", g.Config.Title)
}

// Update handles the game logic update cycle
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw handles the game draw cycle
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.debugDraw(screen)
	g.Level.Draw(screen)
}

func (g *Game) debugDraw(screen *ebiten.Image) {
	if g.Config.DebugObjectLabels {
		ebitenutil.DebugPrint(screen, fmt.Sprint(g))
	}

	if g.Config.DebugGrid {
		drawGrid(screen, g)
	}
}

// Layout defines the game's logical size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Config.SizeX, g.Config.SizeY
}
