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
