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
		DebugInputControlls    bool
	}

	InputState struct {
		CursorPos    [2]int
		FocusedTile  [2]int
		ActiveTile   [2]int
		MouseClicked bool
	}

	Game struct {
		Config  *GameConfig
		InState *InputState

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

	in := &InputState{
		CursorPos:    [2]int{0, 0},
		ActiveTile:   [2]int{-1, -1},
		FocusedTile:  [2]int{0, 0},
		MouseClicked: false,
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
	g := &Game{Config: c, InState: in, Level: c.Levels[0]}

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
	g.updateInputState()
	return nil
}

// Draw handles the game draw cycle
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.debugDraw(screen)
	g.Level.Draw(screen)

	if g.Config.DebugInputControlls {
		drawInputControlls(screen, g)
	}
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

func (g *Game) updateInputState() {
	x, y := ebiten.CursorPosition()

	// Cursor pos
	g.InState.CursorPos[0] = x
	g.InState.CursorPos[1] = y

	// Focused tile
	c := g.Config
	g.InState.FocusedTile[0] = int(x / c.TileSize)
	g.InState.FocusedTile[1] = int(y / c.TileSize)

	// Active buttons
	g.InState.MouseClicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	// Active tile
	if g.InState.MouseClicked {
		g.InState.ActiveTile[0] = g.InState.FocusedTile[0]
		g.InState.ActiveTile[1] = g.InState.FocusedTile[1]
	}
}

func (i *InputState) String() string {
	return fmt.Sprintf("[mouse] x: %d, y: %d\n[ftile] x: %d, y: %d\n[atile] x: %d, y: %d", i.CursorPos[0], i.CursorPos[1], i.FocusedTile[0], i.FocusedTile[1], i.ActiveTile[0], i.ActiveTile[1])
}
