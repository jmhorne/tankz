package projectile

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Cannonball struct {
	x, y float64
}

func NewCannonball(x, y float64) (*Cannonball, error) {
	c := new(Cannonball)
	c.x = x
	c.y = y
	return c, nil
}

func (c *Cannonball) Update() error {
	return nil
}

func (c *Cannonball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(c.x), float32(c.y), 20, color.RGBA{0, 0, 0, 0xff}, true)
}
