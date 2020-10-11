package level

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tjerman/lq/game"
)

const (
	CORRIDOR_DIR_UP = corridorDir(iota)
	CORRIDOR_DIR_RIGHT
	CORRIDOR_DIR_DOWN
	CORRIDOR_DIR_LEFT
)

type (
	corridorDir int
	Room        struct {
		x int
		y int

		// Tilesize
		ts int
	}

	Corridor struct {
		x   int
		y   int
		len int
		dir corridorDir

		// Tilesize
		ts int
	}
)

// determineRooms determines the level's room structure.
// @todo...
func determineRooms(c game.GameConfig) ([]Room, error) {
	cc := make([]Room, 0)
	return cc, nil
}

// determineRoomsF is a debug alternative that provides a constant set
// of rooms.
//
// This should be replaced with a pseudo random world builder-
func determineRoomsF(c *game.GameConfig) ([]*Room, error) {
	cc := make([]*Room, 0)
	ts := c.TileSize

	// cc = append(cc, Room{x: 5, y: 8, ts: ts})
	// cc = append(cc, Room{x: 7, y: 3, ts: ts})

	// cc = append(cc, Room{x: 5, y: 8, ts: ts})
	// cc = append(cc, Room{x: 3, y: 3, ts: ts})

	// cc = append(cc, Room{x: 5, y: 3, ts: ts})
	// cc = append(cc, Room{x: 7, y: 8, ts: ts})

	// cc = append(cc, Room{x: 5, y: 3, ts: ts})
	// cc = append(cc, Room{x: 3, y: 8, ts: ts})

	// cc = append(cc, Room{x: 5, y: 5, ts: ts})
	// cc = append(cc, Room{x: 5, y: 3, ts: ts})

	cc = append(cc, &Room{x: 2, y: 2, ts: ts})
	cc = append(cc, &Room{x: 4, y: 7, ts: ts})
	cc = append(cc, &Room{x: 1, y: 10, ts: ts})
	cc = append(cc, &Room{x: 7, y: 5, ts: ts})
	cc = append(cc, &Room{x: 11, y: 3, ts: ts})
	cc = append(cc, &Room{x: 9, y: 7, ts: ts})

	return cc, nil
}

// determineCorridors connects all of the world's rooms with corredors.
//
// This currently has a basic implementation but will be improved at a later point.
func determineCorridors(c *game.GameConfig, r *Room, rr []*Room) ([]*Corridor, error) {
	cc := make([]*Corridor, 0)

	if len(rr) <= 0 {
		return cc, nil
	}

	// If there is no initial room lets make a pseudo start room, I guess
	invRom := false
	if r == nil {
		r = &Room{x: 0, y: 0}
		invRom = true
	}

	var nr *Room
	nri := -1

	// Find the closest next room to the current room
	var mLen float64 = -1
	for i, cr := range rr {

		// Standard distance equation
		cLen := math.Sqrt(
			math.Pow(float64(r.x)-float64(cr.x), 2) + math.Pow(float64(r.y)-float64(cr.y), 2),
		)

		if mLen < 0 || cLen < mLen {
			mLen = cLen
			nr = cr
			nri = i
		}
	}

	// Remove the next room from the slice so we don't go to infinity
	rrn := make([]*Room, 0, len(rr)-1)
	rrn = append(rrn, rr[0:nri]...)
	rrn = append(rrn, rr[nri+1:]...)

	// Initial rooms don't need a corridor
	if invRom {
		return determineCorridors(c, nr, rrn)
	}

	cc = append(cc, connectRooms(c, r, nr)...)
	ccr, err := determineCorridors(c, nr, rrn)
	if err != nil {
		return nil, err
	}

	return append(cc, ccr...), nil
}

// connectRooms determines a set of corridors that should connect
// the two rooms.
//
// This implementation is... simple... Let's complicate it a bit later.
func connectRooms(c *game.GameConfig, r1, r2 *Room) []*Corridor {
	cc := make([]*Corridor, 0)

	// r2 top right of r1
	if r2.y <= r1.y && r2.x >= r1.x {
		cc = append(cc,
			// From r1 to r2
			&Corridor{
				dir: CORRIDOR_DIR_UP,
				x:   r1.x,
				y:   r1.y,
				len: int(math.Abs(float64(r1.y - r2.y))),
				ts:  c.TileSize,
			},

			// From r2 to r1
			&Corridor{
				dir: CORRIDOR_DIR_LEFT,
				x:   r2.x,
				y:   r2.y,
				len: int(math.Abs(float64(r1.x-r2.x))) - 1,
				ts:  c.TileSize,
			},
		)
	}

	// r2 top left of r1
	if r2.y <= r1.y && r2.x < r1.x {
		cc = append(cc,
			// From r1 to r2
			&Corridor{
				dir: CORRIDOR_DIR_UP,
				x:   r1.x,
				y:   r1.y,
				len: int(math.Abs(float64(r1.y - r2.y))),
				ts:  c.TileSize,
			},

			// From r2 to r1
			&Corridor{
				dir: CORRIDOR_DIR_RIGHT,
				x:   r2.x,
				y:   r2.y,
				len: int(math.Abs(float64(r1.x-r2.x))) - 1,
				ts:  c.TileSize,
			},
		)
	}

	// r2 bottom right of r1
	if r2.y > r1.y && r2.x >= r1.x {
		cc = append(cc,
			// From r1 to r2
			&Corridor{
				dir: CORRIDOR_DIR_DOWN,
				x:   r1.x,
				y:   r1.y,
				len: int(math.Abs(float64(r1.y - r2.y))),
				ts:  c.TileSize,
			},

			// From r2 to r1
			&Corridor{
				dir: CORRIDOR_DIR_LEFT,
				x:   r2.x,
				y:   r2.y,
				len: int(math.Abs(float64(r1.x-r2.x))) - 1,
				ts:  c.TileSize,
			},
		)
	}

	// r2 bottom left of r1
	if r2.y > r1.y && r2.x <= r1.x {
		cc = append(cc,
			// From r1 to r2
			&Corridor{
				dir: CORRIDOR_DIR_DOWN,
				x:   r1.x,
				y:   r1.y,
				len: int(math.Abs(float64(r1.y - r2.y))),
				ts:  c.TileSize,
			},

			// From r2 to r1
			&Corridor{
				dir: CORRIDOR_DIR_RIGHT,
				x:   r2.x,
				y:   r2.y,
				len: int(math.Abs(float64(r1.x-r2.x))) - 1,
				ts:  c.TileSize,
			},
		)
	}

	// Lets remove any empty corridors
	ccClean := make([]*Corridor, 0, len(cc))
	for _, c := range cc {
		if c.len > 0 {
			ccClean = append(ccClean, c)
		}
	}

	return ccClean
}

func drawRooms(s *ebiten.Image, rr []*Room) {
	for i, r := range rr {
		b, _ := ebiten.NewImage(r.ts, r.ts, ebiten.FilterDefault)
		b.Fill(color.White)
		ebitenutil.DebugPrint(b, fmt.Sprintf("room %d", i))
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(r.x*r.ts), float64(r.y*r.ts))
		s.DrawImage(b, op)
	}
}

func drawCorridors(s *ebiten.Image, cc []*Corridor) {
	for i, c := range cc {
		var b *ebiten.Image
		op := &ebiten.DrawImageOptions{}

		if c.dir == CORRIDOR_DIR_UP {
			b, _ = ebiten.NewImage(c.ts, c.ts*c.len, ebiten.FilterDefault)
			op.GeoM.Translate(float64(c.x*c.ts), float64((c.y-c.len)*c.ts))
		}
		if c.dir == CORRIDOR_DIR_DOWN {
			b, _ = ebiten.NewImage(c.ts, c.ts*c.len, ebiten.FilterDefault)
			op.GeoM.Translate(float64(c.x*c.ts), float64((c.y+1)*c.ts))
		}

		if c.dir == CORRIDOR_DIR_LEFT {
			b, _ = ebiten.NewImage(c.ts*c.len, c.ts, ebiten.FilterDefault)
			op.GeoM.Translate(float64((c.x-c.len)*c.ts), float64(c.y*c.ts))
		}
		if c.dir == CORRIDOR_DIR_RIGHT {
			b, _ = ebiten.NewImage(c.ts*c.len, c.ts, ebiten.FilterDefault)
			op.GeoM.Translate(float64((c.x+1)*c.ts), float64(c.y*c.ts))
		}

		b.Fill(color.RGBA{0xff, 0x00, 0x00, 0x90})
		ebitenutil.DebugPrint(b, fmt.Sprintf("cr %d", i))

		s.DrawImage(b, op)
	}
}
