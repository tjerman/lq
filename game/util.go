package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// DrawGrid draws the main world grid on both axises
func drawGrid(s *ebiten.Image, g *Game) {
	c := g.Config
	cc := []color.Color{color.RGBA{0xEF, 0x6F, 0x6C, 0xA0}, color.RGBA{0x56, 0xE3, 0x9F, 0xA0}}

	b, _ := ebiten.NewImage(c.TileSize, c.SizeY, ebiten.FilterDefault)
	for x := 0; x < c.LenX; x++ {
		b.Fill(cc[x%len(cc)])

		if c.DebugObjectLabels {
			ebitenutil.DebugPrint(b, fmt.Sprintf("x: %d", x))
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x*c.TileSize), 0)
		s.DrawImage(b, op)
	}

	b, _ = ebiten.NewImage(c.SizeX, c.TileSize, ebiten.FilterDefault)
	for y := 0; y < c.LenY; y++ {
		b.Fill(cc[y%len(cc)])

		if c.DebugObjectLabels {
			ebitenutil.DebugPrint(b, fmt.Sprintf("y: %d", y))
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(y*c.TileSize))
		s.DrawImage(b, op)
	}
}

func drawInputControlls(s *ebiten.Image, g *Game) {
	dwW := 200
	dwH := 50

	// Draw the debug text
	b, _ := ebiten.NewImage(dwW, dwH, ebiten.FilterDefault)
	b.Fill(color.RGBA{0x18, 0x18, 0x18, 0xA0})
	ebitenutil.DebugPrint(b, fmt.Sprintf("[mouse pos] %s", g.InState))
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.Config.SizeX-dwW), float64(g.Config.SizeY-dwH))
	s.DrawImage(b, op)

	// Draw focused tile
	c := g.Config
	in := g.InState

	b, _ = ebiten.NewImage(c.TileSize, c.TileSize, ebiten.FilterDefault)
	b.Fill(color.RGBA{0x18, 0x18, 0x18, 0xA0})
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(in.FocusedTile[0]*c.TileSize), float64(in.FocusedTile[1]*c.TileSize))
	s.DrawImage(b, op)

	// Draw the active tile
	if in.ActiveTile[0] >= 0 {
		b, _ = ebiten.NewImage(c.TileSize, c.TileSize, ebiten.FilterDefault)
		b.Fill(color.RGBA{0x05, 0x18, 0x05, 0xA0})
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(in.ActiveTile[0]*c.TileSize), float64(in.ActiveTile[1]*c.TileSize))
		s.DrawImage(b, op)
	}
}
